// @Title  
// @Description  
// @Author  yang  2020/7/12 18:32
// @Update  yang  2020/7/12 18:32
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/mail"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// CA 证书颁发机构 (流程图详见：CA流程图 -- CA流程图.png)
/*
		1， CA 生成自签证书
		2,  CA 为别人签发证书
 */
/*

	安装openssl-win64 工具
	1, 修改配置文件  /bin/cnf/openssl.cnf
	2, 终端中执行命令生成私钥 （使用 rsa 算法生成私钥）
		genrsa -out private/cakey.pem 2048
	3, 根据私钥生成自签证书 （详见 CA流程图 -- CA生成自签证书.png）
		req -new -x509 -key private/cakey.pem -out cacert.pem  (根据提示 输入国家，地区，组织，email等信息)

	注意：serial 是序号自增，里面的数字会 自增 (每当新创建一个证书的时候)

	4，在 demoCA 目录下 创建 certs, newcerts, crl 文件夹，创建 index.txt 和 serial文件，并向 serial 文件中写入 01 作为第一张证书的编号。
			详见 CA流程图 -- CA目录结构.png

		### cacert.pem  即为 .pem 格式的证书



	5，为别人签发证书，首先在 PEM 目录下创建 httpd(名称任意), 在 httpd 下创建ssl (名称任意), 创建私钥。
		目录创建完毕，进入 ssl 目录下，CMD 命令 openssl :  (详见  CA流程图 -- CA为别人签发证书.png)

			genrsa -out httpd.key 1024

	6, 填写相应信息，准备申请证书.  (详见 CA流程图 -- CA注册信息.png)

			req -new -key httpd.key -out httpd.csr

	7, 签发证书  (详见 CA流程图 -- CA签发证书.png)

		回到 PEM 路径下,  输入命令：openssl
			ca -in xxx/xxx/httpd.csr -out xxx/xxx/httpd.crt -days 3650

			输入 y 确认签发证书


		### httpd.crt  即为.crt 格式的证书
 */

var errorLog = log.New(os.Stdout, "ERROR", log.Lshortfile)
// 设定一个序列号的最大值
var serialNumberLimit = new(big.Int).Lsh(big.NewInt(1), 128)

func main() {
	// command 表示用户输入的命令
	var command string
	// 用户没有输入任何命令
	if len(os.Args) < 1 {
		command = ""
	}else {
		command = os.Args[1]
	}

	switch command {
	// 证书颁发机构
	case "ca":
		fmt.Println("---------------ca")
		createCA(os.Args[2:], "ca", "/CN=certshop-ca/O=xiongdilian/OU=qukuailian", 3650)
	// 中级证书颁发机构
	case "ica":
		fmt.Println("---------------ica")
		createCA(os.Args[2:], "ca/ica", "/CN=certshop-ca1/O=xiongdilian1/OU=qukuailian1", 365*5)
	// 服务器证书
	case "server":
		fmt.Println("---------------server")
		createCertificate(os.Args[2:], "ca/server", "/CN=server", "127.0.0.1,192.168.12.16,13453234@qq.com",
			365*5, x509.KeyUsageDigitalSignature|x509.KeyUsageDataEncipherment, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth})
	// 客户端证书
	case "client":
		fmt.Println("---------------client")
		createCertificate(os.Args[2:], "ca/client", "/CN=client", "",
			365*5, x509.KeyUsageDigitalSignature|x509.KeyUsageDataEncipherment, []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth})
	default:
		errorLog.Println("Usage: ca | ica | server | client")
	}
}

/*
	1, 如果创建的是CA， 则需要保存其自签证书和私钥，并将其证书拷贝到 ca.pem 中
	2, 如果创建的是 中级CA，则需要将其父级 CA证书拷贝到 自己证书的后面，且将父级CA 中的 ca.pem 拷贝到自己的文件夹中
 */
func createCA(args []string, path, defaultDn string, defaultValidity int) {
	// 参数1：名称     参数2：错误处理策略
	fs := flag.NewFlagSet("ca", flag.PanicOnError)
	dn := fs.String("dn", defaultDn, "证书主题")
	maxPathLength := fs.Int("maxPathLength", 5, "可以颁发证书的最大数量")
	validity := fs.Int("validity", defaultValidity, "证书有效期")
	overwrite := fs.Bool("overwrite", false, "是否覆盖原文件")

	// 解析参数
	err := fs.Parse(args)
	if err != nil {
		errorLog.Fatalf("命令解析失败：%s\n", err)
	}

	// 解析之后剩余的参数 > 1
	if len(fs.Args()) > 1 {
		errorLog.Fatalf("参数非法：%s\n", strings.Join(fs.Args(), ","))
	}else if len(fs.Args()) == 1 {  // 解析之后剩余的参数 = 1
		path = fs.Arg(0)   // 将最后一个有效参数作为 path
	}

	if *overwrite {
		checkExisting(path)
	}

	// 示例1： C:/ca/ica/aa.txt ==>  C:/ca/ica
	// 示例2： ca ==>  .
	ca := filepath.Dir(path)
	var caCert *x509.Certificate
	var caKey *ecdsa.PrivateKey

	// 代表是中级证书
	if ca != "." {
		// 从文件中读取父级 CA 的证书
		caCert = parseCert(ca)
		// 判断是否是 CA
		if !caCert.IsCA {
			errorLog.Fatalf("%s 不是一个有效的证书颁发机构!", ca)
		}else if !(caCert.MaxPathLen > 0) {
			errorLog.Fatalf("证书颁发机构 %s 无法颁发证书！", ca)
		}

		*maxPathLength = caCert.MaxPathLen - 1
		caKey = parseKey(ca)
	}

	key, derKey, err := generatePrivateKey()
	if err != nil {
		errorLog.Fatalf("私钥生成失败！\n")
	}

	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		errorLog.Fatalf("证书编号生成失败: %s\n", err)
	}

	notBefore := time.Now().UTC()
	notAfter := notBefore.AddDate(0,0,*validity)
	template := x509.Certificate{
		// 序列号
		SerialNumber:                serialNumber,
		// 主题
		Subject:*parseDn(caCert, *dn),
		// 生效时间
		NotBefore:notBefore,
		// 失效时间
		NotAfter:notAfter,
		BasicConstraintsValid:true,
		// 是否是CA
		IsCA:true,
		// 可以颁发证书的最大数量
		MaxPathLen:*maxPathLength,
		MaxPathLenZero:*maxPathLength == 0,
		// 证书的用途: 用于数字签名和证书签发
		KeyUsage:x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	// 判断是否获取到父级证书
	if caCert == nil {
		caCert = &template
		caKey = key
	}
	// 关于 x509.CreateCertificate 中是使用 caKey, 还是 key, 要取决于是否存在 根证书CA，
	// 如果不存在根证书，则使用生成 key ，当做根证书的key, 然后使用这个 key, 看下面的代码 （上面有调用）
	/*
			// 判断是否获取到父级证书
			if caCert == nil {
				caCert = &template
				caKey = key
			}
	 */
	// 如果 根证书存在，则使用 根证书的 key.
	derCert, err := x509.CreateCertificate(rand.Reader, &template, caCert, &key.PublicKey, caKey)
	if err != nil {
		errorLog.Fatalf("CA证书创建失败 :%s\n", err)
	}
	// 保存证书
	saveCert(path, derCert)
	// 保存私钥
	saveKey(path, derKey)

	if caCert != &template {
		// 中级CA
		copyFile(filepath.Join(filepath.Dir(path), "ca.pem"),filepath.Join(path, "ca.pem"))
	}else {
		// 最大的 CA
		copyFile(filepath.Join(path,path+".crt"), filepath.Join(path, "ca.pem"))
	}
}

// 拷贝文件
func copyFile(source string, dest string) {
	// 打开源文件
	sourceFile, err := os.Open(source)
	if err != nil {
		errorLog.Fatalf("源文件 %s 打开失败: err : %s\n", source, err)
	}

	defer sourceFile.Close()
	// 打开目标文件
	destFile, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		errorLog.Fatalf("目标文件 %s 打开失败: err : %s\n", dest, err)
	}

	defer destFile.Close()
	if _,err := io.Copy(destFile, sourceFile); err != nil {
		errorLog.Fatalf("拷贝文件 %s 失败：%s\n", source, err )
	}
}


// 保存私钥
func saveKey(dir string ,derKey []byte) {
	fileName := filepath.Join(dir, filepath.Base(dir)+".key")
	keyFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY,  0777)
	if err != nil {
		errorLog.Fatalf("文件 %s 打开失败:%s\n", fileName, err)
	}
	defer keyFile.Close()

	block := &pem.Block{
		Type:    "EC PRIVATE KEY",
		Bytes:   derKey,
	}

	if err := pem.Encode(keyFile, block); err != nil {
		errorLog.Fatalf("私钥 %s 序列化失败:%s\n", fileName,err)
	}
}

// 保存证书
// 如果保存的是 中级证书
func saveCert(dir string, derCert []byte) {
	// 创建文件夹
	createDirectory(dir)
	// ca  ==>   ca/ca.crt
	fileName := filepath.Join(dir, filepath.Base(dir)+".crt")
	certFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		errorLog.Fatalf("文件 %s 打开失败: %s\n", fileName, err)
	}

	defer certFile.Close()

	block := &pem.Block{
		Type:    "CERTIFICATE",
		Bytes:   derCert,
	}
	// 将证书内容序列化至 文件中
	if err := pem.Encode(certFile, block); err != nil {
		errorLog.Fatalf("证书 %s 序列化失败: %s\n", fileName, err)
	}

	// 如果是 中级证书， 需要将父级证书 crt 里面的内容，追加到中级证书的后面，最终生成 中级证书。
	if filepath.Dir(dir) != "." {
		// ca/ica  ca/ca
		newPath := filepath.Join(filepath.Dir(dir), filepath.Base(filepath.Dir(dir)))
		// 打开父级证书
		caFile, err := os.Open(newPath+".crt")
		if err != nil {
			errorLog.Fatalf("证书打开失败: %s\n", err)
		}
		defer  caFile.Close()
		// 拷贝父级证书的内容 至  中级证书
		_,err = io.Copy(certFile, caFile)
		if err != nil {
			errorLog.Fatalf("证书拷贝失败:%s", err)
		}
		// 刷新
		err = certFile.Sync()
		if err != nil {
			errorLog.Fatalf("刷新失败: %s\n", err)
		}
	}
}

// 创建文件夹
func createDirectory(directory string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err := os.MkdirAll(directory, 0777);err != nil  {
			errorLog.Fatalf("文件夹 ./%s 创建失败 : %s\n", directory,err)
		}
	}
}

// 解析用户输入的证书主题
func parseDn(ca *x509.Certificate, dn string) *pkix.Name{
	var caName pkix.Name
	if ca != nil {
		caName = ca.Subject
	}else {
		caName = pkix.Name{}
	}

	newName := &pkix.Name{}
	// -dn="/CN=China/O=xiongdilian/OU=qukuailian"  ===>   CN=China/O=xiongdilian/OU=qukuailian  ===>  CN=China    O=xiongdilian    OU=qukuailian
	for _,element := range strings.Split(strings.Trim(dn, "/"), "/") {
		value := strings.Split(element, "=")
		if len(value) !=2 {
			errorLog.Fatalf("参数 %s 非法!", element)
		}
		switch strings.ToUpper(value[0]) {
		// 名称
		case "CN":
			newName.CommonName = value[1]
		// 国家名称
		case "C":
			if value[1] == "" {
				caName.Country = []string{}
			}else {
				newName.Country = append(newName.Country, value[1])
			}
		// 地点
		case "L":
			if value[1] == "" {
				caName.Locality = []string{}
			}else {
				newName.Locality = append(newName.Locality, value[1])
			}
		// 州或省
		case "ST":
			if value[1] == "" {
				caName.Province = []string{}
			}else {
				newName.Province = append(newName.Province, value[1])
			}
		// 组织
		case "O":
			if value[1] == "" {
				caName.Organization = []string{}
			}else {
				newName.Organization = append(newName.Organization, value[1])
			}
		// 部门
		case "OU":
			if value[1] == "" {
				caName.OrganizationalUnit = []string{}
			}else {
				newName.OrganizationalUnit = append(newName.OrganizationalUnit, value[1])
			}
		default :
			errorLog.Fatalf("参数 %s 非法", element)
		}
	}

	// 将主题 信息 拼接在父级证书的主题 的后面，生成一个新的主题
	if ca != nil {
		newName.Country = append(caName.Country, newName.Country...)
		newName.Locality = append(caName.Locality, newName.Locality...)
		newName.Province = append(caName.Province, newName.Province...)
		newName.Organization = append(caName.Organization, newName.Organization...)
		newName.OrganizationalUnit = append(caName.OrganizationalUnit, newName.OrganizationalUnit...)
	}
	return newName
}

// 生成私钥
// 返回值1: 私钥对象
// 返回值2: 序列化后的私钥, 用于写入文件
func generatePrivateKey() (*ecdsa.PrivateKey, []byte, error){
	// 生成秘钥对
	key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil,nil, err
	}

	// 将私钥序列化
	derKey, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil,nil, err
	}

	return key, derKey, nil
}

// 从文件中读取私钥
func parseKey(path string) *ecdsa.PrivateKey {
	newPath := filepath.Join(path, filepath.Base(path)+".key")
	der, err := ioutil.ReadFile(newPath)
	if err != nil {
		errorLog.Fatalf("私钥文件 %s 读取失败: %s\n", newPath, err)
	}

	// 反序列化
	block, _ := pem.Decode(der)
	if block == nil || block.Type != "EC PRIVATE KEY"{
		errorLog.Fatalf("私钥文件 %s 编码失败: %s \n", newPath, err)
	}

	// 从块儿中解析私钥文件
	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		errorLog.Fatalf("私钥文件 %s 解析失败: %s \n", newPath, err)
	}
	return key
}

// 从文件中读取证书
func parseCert(path string) *x509.Certificate {
	// ca  ==>   /ca/ca.crt
	newPath := filepath.Join(path, filepath.Base(path)+".crt")
	der, err := ioutil.ReadFile(newPath)
	if err != nil {
		errorLog.Fatalf("证书文件 %s 读取失败: %s\n", newPath, err)
	}

	// 反序列化
	block, _ := pem.Decode(der)
	if block == nil || block.Type != "CERTIFICATE"{
		errorLog.Fatalf("证书文件 %s 编码失败: %s \n", newPath, err)
	}

	// 从块儿中解析证书
	crt, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		errorLog.Fatalf("证书文件 %s 解析失败: %s \n", newPath, err)
	}

	return crt
}

// 检查指定路径下是否存在 证书文件，私钥，父级证书文件
func checkExisting(path string) {
	fullPath := filepath.Join(path, filepath.Base(path))  // /ca/ca

	if _,err := os.Stat(fullPath+".crt"); err == nil {
		errorLog.Fatalf("文件: %s 已经存在！", "./"+fullPath+".crt")
	}

	// 判断私钥是否存在
	if _,err := os.Stat(fullPath+".key");err == nil {
		errorLog.Fatalf("文件:%s 已经存在!", "./"+fullPath+".key")
	}

	// 判断父级证书是否存在
	if _,err := os.Stat(fullPath+"ca.pem");err == nil {
		errorLog.Fatalf("文件:%s 已经存在!", "./"+fullPath+"ca.pem")
	}
}

// 创建证书
/*
	参数1：用户输入的参数
	参数2：保存路径
	参数3：证书主题
	参数4: 主题备用名称
	参数5：证书有效时间
	参数6：证书用途
	参数7：证书额外用途
 */
func createCertificate(args []string, path, defaultDn, defaultSan string, defaultValidity int, keyUsage x509.KeyUsage, extKeyUsage []x509.ExtKeyUsage) {
	// 参数1：名称     参数2：错误处理策略
	fs := flag.NewFlagSet("ca", flag.PanicOnError)
	dn := fs.String("dn", defaultDn, "证书主题")
	san := fs.String("san", defaultSan, "备用主题")
	validity := fs.Int("validity", defaultValidity, "证书有效期")
	overwrite := fs.Bool("overwrite", false, "是否覆盖原文件")

	// 解析参数
	err := fs.Parse(args)
	if err != nil {
		errorLog.Fatalf("命令解析失败：%s\n", err)
	}

	// 解析之后剩余的参数 > 1
	if len(fs.Args()) > 1 {
		errorLog.Fatalf("参数非法：%s\n", strings.Join(fs.Args(), ","))
	}else if len(fs.Args()) == 1 {  // 解析之后剩余的参数 = 1
		path = fs.Arg(0)   // 将最后一个有效参数作为 path
	}

	if *overwrite {
		checkExisting(path)
	}

	// 示例1： C:/ca/ica/aa.txt ==>  C:/ca/ica
	// 示例2： ca ==>  .
	var caCert *x509.Certificate
	var caKey *ecdsa.PrivateKey

	ca := filepath.Dir(path)
	caCert = parseCert(ca)
	// 父级证书不是有效的证书颁发机构
	if !caCert.IsCA {
		errorLog.Fatalf("%s 不是证书颁发机构\n", filepath.Dir(path))
	}
	// 获取父级私钥
	caKey = parseKey(ca)

	// 生成证书私钥
	key,derKey, err := generatePrivateKey()
	if err != nil {
		errorLog.Fatalf("私钥生成失败:%s\n", err)
	}

	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		errorLog.Fatalf("证书编号生成失败: %s\n", err)
	}

	notBefore := time.Now().UTC()
	notAfter := notBefore.AddDate(0,0,*validity)
	template := x509.Certificate{
		// 序列号
		SerialNumber:                serialNumber,
		// 主题
		Subject:*parseDn(caCert, *dn),
		// 生效时间
		NotBefore:notBefore,
		// 失效时间
		NotAfter:notAfter,
		BasicConstraintsValid:true,
		// 是否是CA
		IsCA:true,
		// 证书的用途: 用于数字签名和证书签发
		KeyUsage:keyUsage,
		ExtKeyUsage:extKeyUsage,
		EmailAddresses:[]string{},
		IPAddresses:[]net.IP{},
		DNSNames:[]string{},
	}

	// 解析IP, 邮箱，DNS
	parseSubjectNames(*san, &template)
	// 关于 x509.CreateCertificate 中是使用 caKey, 还是 key, 要取决于是否存在 根证书CA，
	// 如果不存在根证书，则使用生成 key ，当做根证书的key, 然后使用这个 key, 看下面的代码 （上面有调用）
	/*
		// 判断是否获取到父级证书
		if caCert == nil {
			caCert = &template
			caKey = key
		}
	*/
	// 如果 根证书存在，则使用 根证书的 key.
	derCert, err := x509.CreateCertificate(rand.Reader, &template, caCert, &key.PublicKey, caKey)
	if err != nil {
		errorLog.Fatalf("证书文件 %s 创建失败: %s", path, err)
	}

	saveCert(path, derCert)
	saveKey(path, derKey)

	// 将父级的 ca.pem 拷贝到自己的文件夹下
	copyFile(filepath.Join(filepath.Dir(path), "ca.pem"), filepath.Join(path, "ca.pem"))

}
// 解析IP, 邮箱，DNS
func parseSubjectNames(san string, template *x509.Certificate) {
	if san != "" {
		for _,h := range strings.Split(san,",") {
			if ip := net.ParseIP(h);ip != nil {
				template.IPAddresses = append(template.IPAddresses, ip)
			}else if email := parseEmailAddress(h); email != nil {
					template.EmailAddresses = append(template.EmailAddresses, email.Address)
			}else {
				template.DNSNames = append(template.DNSNames, h)
			}
		}
	}
}

func parseEmailAddress(address string) *mail.Address{
	email, err := mail.ParseAddress(address)
	if err == nil && email != nil {
		return email
	}
	return nil
}
