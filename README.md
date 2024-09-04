# XV6 installation

To install XV6 on your local machine, you need to have the following tools installed on your system:

- `git`
- `wget`
- `qemu`
- `gdb`

Then you can clone into public xv6 github repository by:

```shell
git clone https://github.com/mit-pdos/xv6-public.git
```

After that, cd into the xv6-public directory and run the following commands:

```shell
make qemu || make qemu-nox
make qemu-gdb
```

For running this os, run the following command in xv6-public directory:

```shell
chmod +x ./kernel && gdb ./kernel
```

