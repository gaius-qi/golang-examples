package main

import (
	"io"
	"log"
	"net"
	"sync"
)

func main() {
	proxy := NewProxy("127.0.0.1:8000", "r-8vbb8713435280d4.redis.zhangbei.rds.aliyuncs.com:6379")
	proxy.Start()
}

type Proxy struct {
	from string
	to   string
	done chan struct{}
}

func NewProxy(from, to string) *Proxy {
	return &Proxy{
		from: from,
		to:   to,
		done: make(chan struct{}),
	}
}

func (p *Proxy) Start() error {
	log.Println("Starting proxy")
	listener, err := net.Listen("tcp", p.from)
	if err != nil {
		return err
	}

	p.run(listener)
	return nil
}

func (p *Proxy) Stop() {
	log.Println("Stopping proxy")
	if p.done == nil {
		return
	}
	close(p.done)
	p.done = nil
}

func (p *Proxy) run(listener net.Listener) {
	for {
		select {
		case <-p.done:
			return
		default:
			connection, err := listener.Accept()
			if err == nil {
				go p.handle(connection)
			} else {
				log.Println("Error accepting conn")
			}
		}
	}
}

func (p *Proxy) handle(connection net.Conn) {
	log.Println("Handling", connection)
	defer log.Println("Done handling", connection)
	defer connection.Close()
	remote, err := net.Dial("tcp", p.to)
	if err != nil {
		log.Println("Error dialing remote host")
		return
	}
	defer remote.Close()
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go p.copy(remote, connection, wg)
	go p.copy(connection, remote, wg)
	wg.Wait()
}

func (p *Proxy) copy(from, to net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-p.done:
		return
	default:
		if _, err := io.Copy(to, from); err != nil {
			log.Println("Error from copy")
			p.Stop()
			return
		}
	}
}
