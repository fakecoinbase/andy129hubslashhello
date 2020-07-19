package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go_6/registry"
	"path"
	"sync"
	"sync/atomic"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// 大map，记录etcd注册中心所有的服务
type AllServiceInfo struct {
	serviceMap map[string]*registry.Service
}

type RegisterService struct {
	id         clientv3.LeaseID
	service    *registry.Service
	registered bool
	// etcd续期
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
}

// 服务注册结构体
type EtcdRegistry struct {
	options *registry.Options
	client  *clientv3.Client
	// 管道，存放的服务
	serviceCh chan *registry.Service
	// 原子操作
	value atomic.Value
	lock  sync.Mutex
	// 服务注册列表
	registryServiceMap map[string]*RegisterService
}

const (
	// 注册中心支持的最大的服务个数
	MaxServiceNum = 8
)

var (
	etcdRegistry *EtcdRegistry = &EtcdRegistry{
		serviceCh:          make(chan *registry.Service, MaxServiceNum),
		registryServiceMap: make(map[string]*RegisterService, MaxServiceNum),
	}
)

// etcd 插件初始化
func init() {
	allServiceInfo := &AllServiceInfo{
		serviceMap: make(map[string]*registry.Service, MaxServiceNum),
	}
	// 进行了原子操作
	etcdRegistry.value.Store(allServiceInfo)
	// 注册插件
	err := registry.RegisterPlugin(etcdRegistry)
	if err != nil {
		err = fmt.Errorf("init registerPlugin err:%v", err)
	}
	//
	go etcdRegistry.run()
}

func (e *EtcdRegistry) run() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		// 服务端开发好的东西，扔到 serviceCh 管理里
		// 这里不断从管道取东西，若管道有值，则加入map记录k-v
		case service := <-e.serviceCh:
			// 先判断下map里有没有
			registryService, ok := e.registryServiceMap[service.Name]
			if ok {
				// 若有此服务，则加入到节点列表即可
				for _, node := range service.Nodes {
					registryService.service.Nodes = append(registryService.service.Nodes, node)
				}
				registryService.registered = false
				break
			}
			registryService = &RegisterService{
				service: service,
			}
			// 若没有，则加入
			e.registryServiceMap[service.Name] = registryService
		case <-ticker.C:
			// etcd注册中心，每隔10秒，进行缓冲更新
			e.syncServiceFromEtcd()
		default:
			// 每500毫秒进行心跳检测
			e.registerOrKeepAlive()
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// 服务端注册列表，在客户端进行缓冲更新
func (e *EtcdRegistry) syncServiceFromEtcd() {
	var allServiceInfoNew = &AllServiceInfo{
		serviceMap: make(map[string]*registry.Service, MaxServiceNum),
	}
	ctx := context.TODO()
	allServiceInfo := e.value.Load().(*AllServiceInfo)
	for _, service := range allServiceInfo.serviceMap {
		key := e.servicePath(service.Name)
		resp, err := e.client.Get(ctx, key, clientv3.WithPrefix())
		if err != nil {
			allServiceInfoNew.serviceMap[service.Name] = service
			continue
		}
		serviceNew := &registry.Service{
			Name: service.Name,
		}
		for _, kv := range resp.Kvs {
			value := kv.Value
			var tmpService registry.Service
			err = json.Unmarshal(value, &tmpService)
			if err != nil {
				fmt.Printf("unmarshal failed, err:%v value:%s", err, string(value))
				return
			}
			for _, node := range tmpService.Nodes {
				serviceNew.Nodes = append(serviceNew.Nodes, node)
			}
		}
		allServiceInfoNew.serviceMap[serviceNew.Name] = serviceNew
	}
	e.value.Store(allServiceInfoNew)
}

func (e *EtcdRegistry) servicePath(name string) string {
	return path.Join(e.options.RegistryPath, name)
}

func (e *EtcdRegistry) registerOrKeepAlive() {
	for _, registryService := range e.registryServiceMap {
		if registryService.registered {
			// 心跳检测
			e.keepAlive(registryService)
			continue
		}
		// 注册
		err := e.registerService(registryService)
		if err != nil {
			fmt.Println("register err")
		}
	}
}

// 心跳检测
func (e *EtcdRegistry) keepAlive(registryService *RegisterService) {
	select {
	// 从etcd续期里直接读数据，若能读到，则心跳检测成功
	case resp := <-registryService.keepAliveCh:
		if resp == nil {
			registryService.registered = false
			return
		}
	}
	return
}

// 注册
func (e *EtcdRegistry) registerService(registryService *RegisterService) (err error) {
	// etcd客户端，设置续期
	resp, err := e.client.Grant(context.TODO(), e.options.HeartBeat)
	if err != nil {
		return
	}
	registryService.id = resp.ID
	for _, node := range registryService.service.Nodes {
		tmp := &registry.Service{
			Name: registryService.service.Name,
			Nodes: []*registry.Node{
				node,
			},
		}
		// 序列化数据
		data, err := json.Marshal(tmp)
		if err != nil {
			continue
		}
		key := e.serviceNodePath(tmp)
		fmt.Printf("register key:%s\n", key)
		// 将服务注册的k-v放入etcd，绑定续期
		_, err = e.client.Put(context.TODO(), key, string(data), clientv3.WithLease(resp.ID))
		if err != nil {
			continue
		}
		// 延长续期
		ch, err := e.client.KeepAlive(context.TODO(), resp.ID)
		if err != nil {
			continue
		}
		// 不断续
		registryService.keepAliveCh = ch
		registryService.registered = true
	}
	return
}

func (e *EtcdRegistry) serviceNodePath(service *registry.Service) string {
	nodeIP := fmt.Sprintf("%s:%d", service.Nodes[0].IP, service.Nodes[0].Port)
	return path.Join(e.options.RegistryPath, service.Name, nodeIP)
}

func (e *EtcdRegistry) Name() string {
	return "etcd"
}

// etcd初始化
func (e *EtcdRegistry) Init(ctx context.Context, opts ...registry.Option) (err error) {
	e.options = &registry.Options{}
	for _, opt := range opts {
		opt(e.options)
	}
	// 选项设计模式的调用
	e.client, err = clientv3.New(clientv3.Config{
		Endpoints:   e.options.Addrs,
		DialTimeout: e.options.Timeout,
	})
	if err != nil {
		err = fmt.Errorf("init etcd err:%v", err)
		return
	}
	return
}

// 注册
func (e *EtcdRegistry) Register(ctx context.Context, service *registry.Service) (err error) {
	select {
	// 服务注册是将服务放入到serviceCh管道，上面有协程不断取值，若管道有值，则存入到etcd的map里进行注册
	case e.serviceCh <- service:
	default:
		err = fmt.Errorf("register chan is full")
		return
	}
	return
}

func (e *EtcdRegistry) Unregister(ctx context.Context, service *registry.Service) (err error) {
	// etcd删除k-v
	return
}

// 服务发现
func (e *EtcdRegistry) GetService(ctx context.Context,
	name string) (service *registry.Service, err error) {
	// 先读缓存
	service, ok := e.getServiceFromCache(ctx, name)
	if ok {
		// 若缓存有，则结束
		return
	}

	// 又读了一遍
	e.lock.Lock()
	defer e.lock.Unlock()
	service, ok = e.getServiceFromCache(ctx, name)
	if ok {
		return
	}

	// 强制从etcd里拉取数据
	key := e.servicePath(name)
	resp, err := e.client.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		// 有异常结束
		return
	}
	// 取到了服务
	service = &registry.Service{
		Name: name,
	}
	// 取到数据后，进行反序列化，数据存到结构体上
	for _, kv := range resp.Kvs {
		value := kv.Value
		var tmpService registry.Service
		err = json.Unmarshal(value, &tmpService)
		if err != nil {
			return
		}
		for _, node := range tmpService.Nodes {
			service.Nodes = append(service.Nodes, node)
		}
	}
	allServiceInfoOld := e.value.Load().(*AllServiceInfo)
	var allServiceInfoNew = &AllServiceInfo{
		serviceMap: make(map[string]*registry.Service, MaxServiceNum),
	}
	for key, val := range allServiceInfoOld.serviceMap {
		allServiceInfoNew.serviceMap[key] = val
	}
	allServiceInfoNew.serviceMap[name] = service
	e.value.Store(allServiceInfoNew)
	return
}

// 从缓存里取
func (e *EtcdRegistry) getServiceFromCache(ctx context.Context,
	name string) (service *registry.Service, ok bool) {
	allServiceInfo := e.value.Load().(*AllServiceInfo)
	// 读map
	service, ok = allServiceInfo.serviceMap[name]
	return
}
