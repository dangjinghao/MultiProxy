package plugins

import (
	"TheresaProxyV2/register"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"strings"
)

type multiProxy struct {
	logger *logrus.Entry
}

func (p multiProxy) Request(*http.Request) error {
	return nil
}
func (p multiProxy) Response(*http.Response) error {
	return nil
}

// 由于service-worker.js会拦截浏览器请求导致请求无法附带session cookie，所以无法切换，写此中间件以检查referer方式实现切换
func (p multiProxy) checkServiceWorker() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if ctx.Request.RequestURI == "/service-worker.js" {
			referer := ctx.GetHeader("referer")
			usefulUrl, _ := url.Parse(referer)
			newTarget := strings.Split(usefulUrl.Path, "/")[1]
			if !register.ExistDomain(newTarget) {
				return
			}
			err := register.SetTargetDomain(ctx, newTarget)
			if err != nil {
				p.logger.Errorf("修改session出错:%s", err)
			}
		}
	}
}

func init() {
	var p multiProxy
	var targetProperty []string
	p.logger = register.PluginLogger("multiProxy")

	supportSites := strings.Split(register.FlagValue("target"), ",")
	redirect := register.FlagValue("redirect")
	var redirectFlag bool
	if redirect == "true" {
		redirectFlag = true
	}
	for _, v := range supportSites {
		targetProperty = strings.Split(v, "://")
		register.ProxySite(targetProperty[1], &register.SiteProperty{Scheme: targetProperty[0],
			Nickname: targetProperty[2], SiteBehavior: p, AutoRedirect: redirectFlag})
	}
	register.Middleware(p.checkServiceWorker())

}
