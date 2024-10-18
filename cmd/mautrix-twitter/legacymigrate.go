package main

import (
	_ "embed"

	up "go.mau.fi/util/configupgrade"
	"maunium.net/go/mautrix/bridgev2/bridgeconfig"
)

const legacyMigrateRenameTables = `
ALTER TABLE portal RENAME TO portal_old;
ALTER TABLE puppet RENAME TO puppet_old;
ALTER TABLE message RENAME TO message_old;
ALTER TABLE reaction RENAME TO reaction_old;
ALTER TABLE "user" RENAME TO user_old;
`

func migrateLegacyConfig(helper up.Helper) {
	bridgeconfig.CopyToOtherLocation(helper, up.Str, []string{"bridge", "displayname_template"}, []string{"network", "displayname_template"})
	bridgeconfig.CopyToOtherLocation(helper, up.Int, []string{"bridge", "displayname_max_length"}, []string{"network", "displayname_max_length"})
}
