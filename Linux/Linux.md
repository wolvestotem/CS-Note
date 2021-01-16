## 关于非阻塞I/O、多路复用、epoll的杂谈

[关于非阻塞I/O、多路复用、epoll的杂谈](https://www.cnblogs.com/upnote/p/12017212.html)

```C++
epfd = epoll_create(EPOLL_QUEUE_LEN); //创建epoll实例
static struct epoll_event ev;
int client_sock;
ev.events = EPOLLIN | EPOLLPRI | EPOLLERR | EPOLLHUP; //声明感兴趣的事件
ev.data.fd = client_sock; //文件描述符指向一个Socket连接
//添加要监控的文件描述符和事件类型到interest list
//真实环境中，可能需要添加成百上千个这种事件。
int res = epoll_ctl(epfd, EPOLL_CTL_ADD, client_sock, &ev); 
while (1) {
int nfds = epoll_wait(epfd, events, MAX_EVENTS_PER_RUN, TIMEOUT);//获取ready的事件，可能有多个
if (nfds < 0) die("Error in epoll_wait!"); //发生错误，退出程序
for(int i = 0; i < nfds; i++) {   //遍历处理每一个已ready的socket
    int fd = events[i].data.fd; 
    handle_io_on_socket(fd);
}
}
```
