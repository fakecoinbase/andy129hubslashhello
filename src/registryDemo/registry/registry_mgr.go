package registry

import (
	"context"
	"fmt"
	"sync"
)

type PluginMgr struct {
	// 用map维护插件，值是接口类型，管理所有插件
	plugins map[string]Registry
	// 锁
	lock sync.Mutex
}

var (
	pluginMgr = &PluginMgr{
		plugins: make(map[string]Registry),
	}
)

// 注册插件
func (p *PluginMgr) RegisterPlugin(registry Registry) (err error) {
	p.lock.Lock()

	// 防止插件覆盖，进行判断
	_, ok := p.plugins[registry.Name()]
	if ok {
		err = fmt.Errorf("registry plugin exist")
		return
	}

	// 直接进行注册
	p.plugins[registry.Name()] = registry
	p.lock.Unlock()
	return
}

// 进行初始化注册中心
func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return pluginMgr.initRegistry(ctx, name, opts...)
}

func (p *PluginMgr) initRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	// 先查看服务列表，服务是否存在，若不存在，就没的初始化，报错
	plugin, ok := p.plugins[name]
	if !ok {
		err = fmt.Errorf("plugin %s not exist", name)
		return
	}
	// 存在，返回值赋值
	registry = plugin
	// 进行组件初始化
	err = plugin.Init(ctx, opts...)
	return
}