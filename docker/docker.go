package docker

import (
	"docker-api-service/app"
	"docker-api-service/userdatabase"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/strslice"
	"github.com/gin-gonic/gin"
	"os/exec"
)

func CreateDocker(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	form := CreateStruct{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		appG.Response(0, err.Error(), "")
		return
	}
	ip := userdatabase.Okip(form.Ipprefix)
	dockerconfig := &container.Config{
		Hostname: form.Hostname,
		//ExposedPorts: nat.PortSet{"22/tcp": {}},
		Image: form.Image,
		Labels: map[string]string{
			"org.label-schema.tc.enabled": "1",
			"org.label-schema.tc.rate":    form.MinNet, //最低宽带限制
			"org.label-schema.tc.ceil":    form.MaxNet, //最大宽带限制
		},
		Cmd: strslice.StrSlice{"sh", "-c", "/etc/init.d/ssh start && tail -f /dev/null"},
	}
	dockerHostConfig := &container.HostConfig{
		//PortBindings: nat.PortMap{"22/tcp": []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: "41654"}}},
		Resources: container.Resources{
			CPUShares:  form.CPUShares,        //此容器相对于其他容器的相对CPU权重
			CpusetCpus: form.CpusetCpus,       //运行的核心exp:0-3或0,1
			NanoCPUs:   form.NanoCPUs,         //相当于--cpus=2
			Memory:     form.Memory * 1048576, //以字节为单位
			MemorySwap: form.MemorySwap,       //总内存包含swap,-1未无限制使用swap
		},
		StorageOpt: map[string]string{"size": "11G"},
	}
	dockernetconfig := &network.NetworkingConfig{EndpointsConfig: map[string]*network.EndpointSettings{
		"nat": {
			IPAMConfig: &network.EndpointIPAMConfig{IPv4Address: ip}},
	}}
	_, err := cli.ContainerCreate(ctx, dockerconfig, dockerHostConfig, dockernetconfig, nil, form.Hostname)
	if err != nil {
		appG.Response(0, err.Error(), "")
		return
	} else {
		userdatabase.Create(form.Hostname, ip, form.ProtNums, form.DomainNums)
		appG.Response(200, "创建成功", ip)
		return
	}
}
func Dockerinfo(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	id := ctx.PostForm("id")
	if id == "" {
		appG.Response(0, "不是有效的Id标识", "")
		return
	} else {
		info, err := cli.ContainerInspect(ctx, id)
		if err != nil {
			appG.Response(0, err.Error(), "")
			return
		} else {
			if info.Name != "/"+id {
				appG.Response(0, "该用户不存在", "")
				return
			} else {
				appG.Response(200, "获取成功", info)
				return
			}
		}
	}
}
func DeleteDocker(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	id := ctx.PostForm("id")
	if id == "" {
		appG.Response(0, "不是有效的Id标识", "")
		return
	} else {
		info, err := cli.ContainerInspect(ctx, id)
		if err != nil {
			appG.Response(0, "该用户不存在", "")
			return
		} else {
			if info.Name != "/"+id {
				appG.Response(0, "该用户不存在", "")
				return
			} else {
				err = cli.ContainerRemove(ctx, info.ID, types.ContainerRemoveOptions{})
				if err != nil {
					appG.Response(0, err.Error(), "")
					return
				} else {
					appG.Response(200, "删除容器成功", "")
					return
				}
			}
		}
	}
}
func PauseDocker(ctx *gin.Context) { //暂停容器的所有进程
	appG := app.Gin{C: ctx}
	id := ctx.PostForm("id")
	if id == "" {
		appG.Response(0, "不是有效的Id标识", "")
		return
	} else {
		info, err := cli.ContainerInspect(ctx, id)
		if err != nil {
			appG.Response(0, "该用户不存在", "")
			return
		} else {
			if info.Name != "/"+id {
				appG.Response(0, "该用户不存在", "")
				return
			} else {
				err = cli.ContainerPause(ctx, info.ID)
				if err != nil {
					appG.Response(0, err.Error(), "")
					return
				} else {
					appG.Response(200, "冻结进程成功", "")
					return
				}
			}
		}
	}
}
func UnPauseDocker(ctx *gin.Context) { //恢复暂停容器的所有进程
	appG := app.Gin{C: ctx}
	id := ctx.PostForm("id")
	if id == "" {
		appG.Response(0, "不是有效的Id标识", "")
		return
	} else {
		info, err := cli.ContainerInspect(ctx, id)
		if err != nil {
			appG.Response(0, "该用户不存在", "")
			return
		} else {
			if info.Name != "/"+id {
				appG.Response(0, "该用户不存在", "")
				return
			} else {
				err = cli.ContainerUnpause(ctx, info.ID)
				if err != nil {
					appG.Response(0, err.Error(), "")
					return
				} else {
					appG.Response(200, "恢复冻结进程成功", "")
					return
				}
			}
		}
	}
}
func StartDocker(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	id := ctx.PostForm("id")
	if id == "" {
		appG.Response(0, "不是有效的Id标识", "")
		return
	} else {
		info, err := cli.ContainerInspect(ctx, id)
		if err != nil {
			appG.Response(0, "该用户不存在", "")
			return
		} else {
			if info.Name != "/"+id {
				appG.Response(0, "该用户不存在", "")
				return
			} else {
				err = cli.ContainerStart(ctx, info.ID, types.ContainerStartOptions{})
				if err != nil {
					appG.Response(0, err.Error(), "")
					return
				} else {
					appG.Response(200, "启动容器成功", "")
					return
				}
			}
		}
	}
}
func StopDocker(ctx *gin.Context) { //恢复暂停容器的所有进程
	appG := app.Gin{C: ctx}
	id := ctx.PostForm("id")
	if id == "" {
		appG.Response(0, "不是有效的Id标识", "")
		return
	} else {
		info, err := cli.ContainerInspect(ctx, id)
		if err != nil {
			appG.Response(0, "该用户不存在", "")
			return
		} else {
			if info.Name != "/"+id {
				appG.Response(0, "该用户不存在", "")
				return
			} else {
				err = cli.ContainerStop(ctx, info.ID, nil)
				if err != nil {
					appG.Response(0, err.Error(), "")
					return
				} else {
					appG.Response(200, "停止容器成功", "")
					return
				}
			}
		}
	}
}
func RestarDocker(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	id := ctx.PostForm("id")
	if id == "" {
		appG.Response(0, "不是有效的Id标识", "")
		return
	} else {
		info, err := cli.ContainerInspect(ctx, id)
		if err != nil {
			appG.Response(0, "该用户不存在", "")
			return
		} else {
			if info.Name != "/"+id {
				appG.Response(0, "该用户不存在", "")
				return
			} else {
				err = cli.ContainerRestart(ctx, info.ID, nil)
				if err != nil {
					appG.Response(0, err.Error(), "")
					return
				} else {
					appG.Response(200, "重启容器成功", "")
					return
				}
			}
		}
	}
}
func AddPort(ctx *gin.Context) { //恢复暂停容器的所有进程
	appG := app.Gin{C: ctx}
	id := ctx.PostForm("id")
	porto := ctx.PostForm("porto")
	port := ctx.PostForm("port")
	typee := ctx.PostForm("type")
	if id == "" {
		appG.Response(0, "不是有效的Id标识", "")
		return
	} else {
		info, err := cli.ContainerInspect(ctx, id)
		if err != nil {
			appG.Response(0, err.Error(), "")
			return
		} else {
			if info.Name != "/"+id {
				appG.Response(0, "该用户不存在", "")
				return
			} else {
				user := userdatabase.SearchUser(id)
				if iptablesexist(user.Ip, porto, port, typee) == true {
					appG.Response(0, "该规则已存在", "")
					return
				} else {
					if userdatabase.UpdateUserProt(id, porto, port) == true {
						cmd := exec.Command("iptables", "-t", "nat", "-A", "PREROUTING", "-i", "eth0", "-p", typee, "-m", typee, "--dport", porto, "-j", "DNAT", "--to-destination", user.Ip+":"+port)
						err := cmd.Run()
						if err != nil {
							appG.Response(0, "添加端口失败", "")
							return
						} else {
							appG.Response(200, "添加端口成功", "")
							cmd := exec.Command("/bin/bash", "-c", "service iptables save && service iptables restart")
							cmd.Run()
							return
						}
					} else {
						appG.Response(0, "添加端口失败", "")
						return
					}
				}
			}
		}
	}
}
func DelPort(ctx *gin.Context) { //恢复暂停容器的所有进程
	appG := app.Gin{C: ctx}
	id := ctx.PostForm("id")
	porto := ctx.PostForm("porto")
	port := ctx.PostForm("port")
	typee := ctx.PostForm("type")
	if id == "" {
		appG.Response(0, "不是有效的Id标识", "")
		return
	} else {
		info, err := cli.ContainerInspect(ctx, id)
		if err != nil {
			appG.Response(0, err.Error(), "")
			return
		} else {
			if info.Name != "/"+id {
				appG.Response(0, "该用户不存在", "")
				return
			} else {
				user := userdatabase.SearchUser(id)
				cmd := exec.Command("iptables", "-t", "nat", "-D", "PREROUTING", "-i", "eth0", "-p", typee, "-m", typee, "--dport", porto, "-j", "DNAT", "--to-destination", user.Ip+":"+port)
				err := cmd.Run()
				if err != nil {
					appG.Response(0, "删除端口失败", "")
					return
				} else {
					appG.Response(200, "删除端口成功", "")
					userdatabase.DeleteUserProt(id, porto)
					cmd := exec.Command("/bin/bash", "-c", "service iptables save && service iptables restart")
					cmd.Run()
					return
				}
			}
		}
	}
}
func AddDomain(ctx *gin.Context) { //恢复暂停容器的所有进程
	appG := app.Gin{C: ctx}
	id := ctx.PostForm("id")
	domain := ctx.PostForm("domain")
	if id == "" {
		appG.Response(0, "不是有效的Id标识", "")
		return
	} else {
		info, err := cli.ContainerInspect(ctx, id)
		if err != nil {
			appG.Response(0, err.Error(), "")
			return
		} else {
			if info.Name != "/"+id {
				appG.Response(0, "该用户不存在", "")
				return
			} else {
				if userdatabase.UpdateUserDomain(id, domain) == true {
					/*
						cmd := exec.Command("iptables", "-t", "nat", "-A", "PREROUTING", "-i", "eth0", "-p", typee, "-m", typee, "--dport", porto, "-j", "DNAT", "--to-destination", user.Ip+":"+port)
						err := cmd.Run()
						if err != nil {
							appG.Response(0, "添加域名失败", "")
							return
						} else {
							appG.Response(0, "添加域名成功", "")
							cmd := exec.Command("/bin/bash", "-c", "/etc/init.d/caddy restart")
							cmd.Run()
							return
						}
					*/
					appG.Response(200, "添加域名成功", "")
					return
				} else {
					appG.Response(0, "添加域名失败", "")
					return
				}
			}
		}
	}
}
func DelDomain(ctx *gin.Context) { //恢复暂停容器的所有进程
	appG := app.Gin{C: ctx}
	id := ctx.PostForm("id")
	//domain := ctx.PostForm("domain")
	if id == "" {
		appG.Response(0, "不是有效的Id标识", "")
		return
	} else {
		info, err := cli.ContainerInspect(ctx, id)
		if err != nil {
			appG.Response(0, err.Error(), "")
			return
		} else {
			if info.Name != "/"+id {
				appG.Response(0, "该用户不存在", "")
				return
			} else {
				/*
					user := userdatabase.SearchUser(id)
					cmd := exec.Command("iptables", "-t", "nat", "-D", "PREROUTING", "-i", "eth0", "-p", typee, "-m", typee, "--dport", porto, "-j", "DNAT", "--to-destination", user.Ip+":"+port)
					err := cmd.Run()
					if err != nil {
						appG.Response(0, "删除端口失败", "")
						return
					} else {
						appG.Response(0, "删除端口成功", "")
						userdatabase.DeleteUserProt(id, porto)
						cmd := exec.Command("/bin/bash", "-c", "service iptables save && service iptables restart")
						cmd.Run()
						return
					}
				*/
				appG.Response(200, "删除域名成功", "")
				return
			}
		}
	}
}
