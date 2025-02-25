//go:build darwin

package table

import (
	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/table"

	"github.com/macadmins/osquery-extension/tables/filevaultusers"
	"github.com/macadmins/osquery-extension/tables/macos_profiles"
	"github.com/macadmins/osquery-extension/tables/mdm"
	"github.com/macadmins/osquery-extension/tables/munki"
	"github.com/macadmins/osquery-extension/tables/unifiedlog"
)

func platformTables() []osquery.OsqueryPlugin {
	return []osquery.OsqueryPlugin{
		table.NewPlugin("filevault_users", filevaultusers.FileVaultUsersColumns(), filevaultusers.FileVaultUsersGenerate),
		table.NewPlugin("macos_profiles", macos_profiles.MacOSProfilesColumns(), macos_profiles.MacOSProfilesGenerate),
		table.NewPlugin("mdm", mdm.MDMInfoColumns(), mdm.MDMInfoGenerate),
		table.NewPlugin("munki_info", munki.MunkiInfoColumns(), munki.MunkiInfoGenerate),
		table.NewPlugin("munki_installs", munki.MunkiInstallsColumns(), munki.MunkiInstallsGenerate),
		// osquery version 5.5.0 and up ships a unified_log table in core
		// we are renaming the one from the macadmins extension to avoid collision
		table.NewPlugin("macadmins_unified_log", unifiedlog.UnifiedLogColumns(), unifiedlog.UnifiedLogGenerate),
	}
}
