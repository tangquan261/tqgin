package tokenbuilder

import (
	"tqgin/pkg/AccessTocken"
)

// Role Type
type RoleRTM uint16

// Role consts
const (
	RoleRtmUser = 1
)

//RtmTokenBuilder class
type RtmTokenBuilder struct {
}

//BuildToken method
// appID: The App ID issued to you by Agora. Apply for a new App ID from
//        Agora Dashboard if it is missing from your kit. See Get an App ID.
// appCertificate:	Certificate of the application that you registered in
//                  the Agora Dashboard. See Get an App Certificate.
// userAccount: The user account.
// role: Role_Rtm_User = 1
// privilegeExpireTs: represented by the number of seconds elapsed since
//                    1/1/1970. If, for example, you want to access the
//                    Agora Service within 10 minutes after the token is
//                    generated, set expireTimestamp as the current
//                    timestamp + 600 (seconds)./
func RTMBuildToken(appID string, appCertificate string, userAccount string, role RoleRTM, privilegeExpiredTs uint32) (string, error) {

	token := accesstoken.CreateAccessToken2(appID, appCertificate, userAccount, "")
	token.AddPrivilege(accesstoken.KLoginRtm, privilegeExpiredTs)

	return token.Build()
}
