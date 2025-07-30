package network

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/Laisky/zap"

	"github.com/songquanpeng/one-api/common/logger"
)

func splitSubnets(subnets string) []string {
	res := strings.Split(subnets, ",")
	for i := 0; i < len(res); i++ {
		res[i] = strings.TrimSpace(res[i])
	}
	return res
}

func isValidSubnet(subnet string) error {
	_, _, err := net.ParseCIDR(subnet)
	if err != nil {
		return fmt.Errorf("failed to parse subnet: %w", err)
	}
	return nil
}

func isIpInSubnet(ctx context.Context, ip string, subnet string) bool {
	_, ipNet, err := net.ParseCIDR(subnet)
	if err != nil {
		logger.Logger.Error("failed to parse subnet", zap.String("subnet", subnet), zap.Error(err))
		return false
	}
	return ipNet.Contains(net.ParseIP(ip))
}

func IsValidSubnets(subnets string) error {
	for _, subnet := range splitSubnets(subnets) {
		if err := isValidSubnet(subnet); err != nil {
			return err
		}
	}
	return nil
}

func IsIpInSubnets(ctx context.Context, ip string, subnets string) bool {
	for _, subnet := range splitSubnets(subnets) {
		if isIpInSubnet(ctx, ip, subnet) {
			return true
		}
	}
	return false
}
