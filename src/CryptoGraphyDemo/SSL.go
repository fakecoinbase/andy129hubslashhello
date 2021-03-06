// @Title  
// @Description  
// @Author  yang  2020/7/12 17:06
// @Update  yang  2020/7/12 17:06
package main

// SSL 协议
/*
	1, 简介
		应用层很多传输协议是不安全的，比如 http, SMTP, ftp,这些协议在互联网上以明文的形式传输数据，
		在互联网早期，网景公司（netscape） 在应用层和传输层加了半个层，主要用于保密和鉴别，后来国际标准化组织研发了
		更加流行的通用协议 TLS。
	2, SSL 协议的特性
		a, 保密：在握手协议中定义了会话秘钥，所有消息都会被加密
		b, 鉴别：可选的客户端认证 和 强制的服务器端验证
		c, 完整性：传送的消息包括消息完整性检查 (使用 MAC )
		MAC：含有秘钥的散列函数，兼容 MD 和 SHA 算法的特性，并在此基础加入 秘钥，是一种更加安全的消息摘要算法。
	3，SSL协议的工作原理
		SSL 协议主要是有 三个自协议组成，分别是 握手协议，记录协议，警报协议。
			3.1 握手协议
				--- 握手协议是客户端和服务器端 用 SSL 链接通信时使用的第一个子协议，包括客户度与服务器之间的一系列消息，
					主要是协商加密，哈希算法，签名算法，这些算法主要用来保护传输数据。
					主要包含三个字段，分别是 Type: 消息类型，Length:消息长度，Content:消息内容
				握手协议 又分四个阶段：
					1, 建立安全能力， SSL 握手的第一阶段启动逻辑连接。 (详见 SSL -- SSL建立安全能力流程图.png)
						首先客户端向服务器端发送 client hello 消息并等待服务器响应，随后服务器向客户端 发送 server hello,
						对client hello 进行确认。

						ClientHello 主要包含的信息：
						1, 客户端可以支持的 SSL 最高版本
						2, 一个用于生成主秘钥的 32字节的随机数
						3, 一个客户端可以支持的密码算法列表
						密码套件的格式：每个套件以 SSL 开头。
						SSL_DHE_RSA_WITH_DES_CBC_WITH_SHA

						ServerHello 主要包含的信息：
						1, 服务器端所采用的 SSL 版本号
						2, 一个用于生成主秘钥的 32字节的随机数

					2, 服务器鉴别与秘钥交换  （详见 SSL -- SSL服务器端鉴别与秘钥交换.png）
						服务器启动 SSL 握手第2 阶段，是消息的唯一发送方，客户端是消息的唯一接收方。

					3, 客户端鉴别与秘钥交换
						客户端启动SSL 握手协议的第三阶段， 是本阶段所有消息的唯一发送方，服务器是所有消息唯一的接收方。
						1，证书：为了对服务器证明自己，客户端要发送一个证书信息
						2，客户端发送主秘钥给服务端，会使用服务器端的公钥进行加密。
						3，对服务器端进行验证。

					4，完成阶段
						客户端启动 SSL 握手第四阶段，结束会话。

			3.2 记录协议		（详见 SSL -- SSL记录协议.png）
				--- 记录协议 在客户端和服务端握手成功后使用，进入 SSL 记录协议， 主要提供两个服务：
					1，保密性： 使用握手协议商议的秘钥进行实现
					2，完整性： 握手协议定义了MAC，用于保证消息的完整性

			3.3 警报协议
				--- 客户端和服务端如果发现有错误，会向对方发送一个警报信息，如果是 致命错误，则会立即关闭 SSL 链接，双方会删除相关的会话ID，秘钥等等。
					每个警报共2个字节，第一个字节表示错误类型，如果是警报则值为1， 如果是致命错误，则值为2；  第二个字节是具体的实际错误。
 */

func main() {
	
}
