package headupdater

import "github.com/Kimenzo/any-sync/commonspace/object/acl/list"

type AclUpdater interface {
	UpdateAcl(aclList list.AclList)
}
