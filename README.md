GoVdpa contains a set of libraries and programs to manage vdpa devices in Golang

### Libraries: 
- **uvdpa**: Implements the interaction with Userspace Vdpa Daemon

### Programs:
- **uvpda-cli**: A command line interface to the Userspace Vdpa Daemon

# uvdpa-cli
uvdpa-cli is a command line interface that can send commands to the Userspace Vdpa Daemon.

## Building
In order to build the application, just run the top level Makefile:

    $ make

## Running uvdpa-cli
If the program is executed without arguments, you will enter in *interactive mode*.

    ./build/uvdpa-cli
    Staring userspace vdpa cli (type "help" to list the available commands)
    [vdpacli] $ help
    Commands
    help                               Print this help
    exit                               Quit program
    version                            Print daemon version
    list                               List configured interfaces
    create [DEVICE] [SOCKET] [MODE]    Create interface
                DEVICE: A device PCI address, e.g: 00:0f:01
                SOCKET: A socketfile path, e.g: /tmp/vdpa1.sock
                MODE: [client|server]
    destroy [DEVICE]                   Destroy a device
                DEVICE: A device PCI address, e.g: 00:0f:01
    [vdpacli] $ 


Extra arguments will be interpreted as commands (as if they were introduced in the *interactive mode* shell). After trying to run the specified commands, the program will exit, e.g:

    $ ./build/uvdpa-cli list
    device: 000:05:00.3      socket: /tmp/vdpa1      mode: server    driver: net_ifcvf

    $ ./build/uvdpa-cli create 000:05:00.4 /tmp/vdpa2 client && ./build/vdpacli list
    Success
    device: 000:05:00.3      socket: /tmp/vdpa1      mode: server    driver: net_ifcvf
    device: 000:05:00.4      socket: /tmp/vdpa2      mode: client    driver: net_ifcvf
    $

