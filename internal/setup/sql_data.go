package setup

import (
	"github.com/TeaOSLab/EdgeAPI/internal/acme"
	"github.com/TeaOSLab/EdgeAPI/internal/db/models"
	"github.com/TeaOSLab/EdgeAPI/internal/errors"
	"github.com/TeaOSLab/EdgeAPI/internal/utils"
	"github.com/iwind/TeaGo/dbs"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/rands"
	"github.com/iwind/TeaGo/types"
	stringutil "github.com/iwind/TeaGo/utils/string"
)

type upgradeVersion struct {
	version string
	f       func(db *dbs.DB) error
}

var upgradeFuncs = []*upgradeVersion{
	{
		"0.0.3", upgradeV0_0_3,
	},
	{
		"0.0.5", upgradeV0_0_5,
	},
	{
		"0.0.6", upgradeV0_0_6,
	},
	{
		"0.0.9", upgradeV0_0_9,
	},
	{
		"0.0.10", upgradeV0_0_10,
	},
}

// 升级SQL数据
func UpgradeSQLData(db *dbs.DB) error {
	version, err := db.FindCol(0, "SELECT version FROM edgeVersions")
	if err != nil {
		return err
	}
	versionString := types.String(version)
	if len(versionString) > 0 {
		for _, f := range upgradeFuncs {
			if stringutil.VersionCompare(versionString, f.version) >= 0 {
				continue
			}
			err = f.f(db)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// v0.0.3
func upgradeV0_0_3(db *dbs.DB) error {
	// 获取第一个管理员
	adminIdCol, err := db.FindCol(0, "SELECT id FROM edgeAdmins ORDER BY id ASC LIMIT 1")
	if err != nil {
		return err
	}
	adminId := types.Int64(adminIdCol)
	if adminId <= 0 {
		return errors.New("'edgeAdmins' table should not be empty")
	}

	// 升级edgeDNSProviders
	_, err = db.Exec("UPDATE edgeDNSProviders SET adminId=? WHERE adminId=0 AND userId=0", adminId)
	if err != nil {
		return err
	}

	// 升级edgeDNSDomains
	_, err = db.Exec("UPDATE edgeDNSDomains SET adminId=? WHERE adminId=0 AND userId=0", adminId)
	if err != nil {
		return err
	}

	// 升级edgeSSLCerts
	_, err = db.Exec("UPDATE edgeSSLCerts SET adminId=? WHERE adminId=0 AND userId=0", adminId)
	if err != nil {
		return err
	}

	// 升级edgeNodeClusters
	_, err = db.Exec("UPDATE edgeNodeClusters SET adminId=? WHERE adminId=0 AND userId=0", adminId)
	if err != nil {
		return err
	}

	// 升级edgeNodes
	_, err = db.Exec("UPDATE edgeNodes SET adminId=? WHERE adminId=0 AND userId=0", adminId)
	if err != nil {
		return err
	}

	// 升级edgeNodeGrants
	_, err = db.Exec("UPDATE edgeNodeGrants SET adminId=? WHERE adminId=0", adminId)
	if err != nil {
		return err
	}

	return nil
}

// v0.0.5
func upgradeV0_0_5(db *dbs.DB) error {
	// 升级edgeACMETasks
	_, err := db.Exec("UPDATE edgeACMETasks SET authType=? WHERE authType IS NULL OR LENGTH(authType)=0", acme.AuthTypeDNS)
	if err != nil {
		return err
	}

	return nil
}

// v0.0.6
func upgradeV0_0_6(db *dbs.DB) error {
	stmt, err := db.Prepare("SELECT COUNT(*) FROM edgeAPITokens WHERE role='user'")
	if err != nil {
		return err
	}
	defer func() {
		_ = stmt.Close()
	}()
	col, err := stmt.FindCol(0)
	if err != nil {
		return err
	}
	count := types.Int(col)
	if count > 0 {
		return nil
	}

	nodeId := rands.HexString(32)
	secret := rands.String(32)
	_, err = db.Exec("INSERT INTO edgeAPITokens (nodeId, secret, role) VALUES (?, ?, ?)", nodeId, secret, "user")
	if err != nil {
		return err
	}

	return nil
}

// v0.0.9
func upgradeV0_0_9(db *dbs.DB) error {
	// firewall policies
	var tx *dbs.Tx
	dbs.NotifyReady()
	policies, err := models.NewHTTPFirewallPolicyDAO().FindAllEnabledFirewallPolicies(tx)
	if err != nil {
		return err
	}
	for _, policy := range policies {
		if policy.ServerId > 0 {
			continue
		}
		policyId := int64(policy.Id)
		webIds, err := models.NewHTTPWebDAO().FindAllWebIdsWithHTTPFirewallPolicyId(tx, policyId)
		if err != nil {
			return err
		}
		serverIds := []int64{}
		for _, webId := range webIds {
			serverId, err := models.NewServerDAO().FindEnabledServerIdWithWebId(tx, webId)
			if err != nil {
				return err
			}
			if serverId > 0 && !lists.ContainsInt64(serverIds, serverId) {
				serverIds = append(serverIds, serverId)
			}
		}
		if len(serverIds) == 1 {
			err = models.NewHTTPFirewallPolicyDAO().UpdateFirewallPolicyServerId(tx, policyId, serverIds[0])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// v0.0.10
func upgradeV0_0_10(db *dbs.DB) error {
	// IP Item列表转换
	ones, _, err := db.FindOnes("SELECT * FROM edgeIPItems ORDER BY id ASC")
	if err != nil {
		return err
	}
	for _, one := range ones {
		ipFromLong := utils.IP2Long(one.GetString("ipFrom"))
		ipToLong := utils.IP2Long(one.GetString("ipTo"))
		_, err = db.Exec("UPDATE edgeIPItems SET ipFromLong=?, ipToLong=? WHERE id=?", ipFromLong, ipToLong, one.GetInt64("id"))
		if err != nil {
			return err
		}
	}

	return nil
}
