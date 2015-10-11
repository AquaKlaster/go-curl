package curl

/*
#include <curl/curl.h>
#include "compat.h"

*/
import "C"

// for GlobalInit(flag)
const (
	GLOBAL_SSL     = C.CURL_GLOBAL_SSL
	GLOBAL_WIN32   = C.CURL_GLOBAL_WIN32
	GLOBAL_ALL     = C.CURL_GLOBAL_ALL
	GLOBAL_NOTHING = C.CURL_GLOBAL_NOTHING
	GLOBAL_DEFAULT = C.CURL_GLOBAL_DEFAULT
)

// CURLMcode
const (
	M_CALL_MULTI_PERFORM = C.CURLM_CALL_MULTI_PERFORM
	//        M_OK                 = C.CURLM_OK
	M_BAD_HANDLE      = C.CURLM_BAD_HANDLE
	M_BAD_EASY_HANDLE = C.CURLM_BAD_EASY_HANDLE
	M_OUT_OF_MEMORY   = C.CURLM_OUT_OF_MEMORY
	M_INTERNAL_ERROR  = C.CURLM_INTERNAL_ERROR
	M_BAD_SOCKET      = C.CURLM_BAD_SOCKET
	M_UNKNOWN_OPTION  = C.CURLM_UNKNOWN_OPTION
)

// for multi.Setopt(flag, ...)
const (
	MOPT_SOCKETFUNCTION = C.CURLMOPT_SOCKETFUNCTION
	MOPT_SOCKETDATA     = C.CURLMOPT_SOCKETDATA
	MOPT_PIPELINING     = C.CURLMOPT_PIPELINING
	MOPT_TIMERFUNCTION  = C.CURLMOPT_TIMERFUNCTION
	MOPT_TIMERDATA      = C.CURLMOPT_TIMERDATA
	MOPT_MAXCONNECTS    = C.CURLMOPT_MAXCONNECTS
)

// CURLSHcode
const (
	//        SHE_OK         = C.CURLSHE_OK
	SHE_BAD_OPTION = C.CURLSHE_BAD_OPTION
	SHE_IN_USE     = C.CURLSHE_IN_USE
	SHE_INVALID    = C.CURLSHE_INVALID
	SHE_NOMEM      = C.CURLSHE_NOMEM
)

// for share.Setopt(flat, ...)
const (
	SHOPT_SHARE      = C.CURLSHOPT_SHARE
	SHOPT_UNSHARE    = C.CURLSHOPT_UNSHARE
	SHOPT_LOCKFUNC   = C.CURLSHOPT_LOCKFUNC
	SHOPT_UNLOCKFUNC = C.CURLSHOPT_UNLOCKFUNC
	SHOPT_USERDATA   = C.CURLSHOPT_USERDATA
)

// for share.Setopt(SHOPT_SHARE/SHOPT_UNSHARE, flag)
const (
	LOCK_DATA_SHARE       = C.CURL_LOCK_DATA_SHARE
	LOCK_DATA_COOKIE      = C.CURL_LOCK_DATA_COOKIE
	LOCK_DATA_DNS         = C.CURL_LOCK_DATA_DNS
	LOCK_DATA_SSL_SESSION = C.CURL_LOCK_DATA_SSL_SESSION
	LOCK_DATA_CONNECT     = C.CURL_LOCK_DATA_CONNECT
)

// for VersionInfo(flag)
const (
	VERSION_FIRST  = C.CURLVERSION_FIRST
	VERSION_SECOND = C.CURLVERSION_SECOND
	VERSION_THIRD  = C.CURLVERSION_THIRD
	// VERSION_FOURTH = C.CURLVERSION_FOURTH
	VERSION_LAST = C.CURLVERSION_LAST
	VERSION_NOW  = C.CURLVERSION_NOW
)

// for VersionInfo(...).Features mask flag
const (
	VERSION_IPV6         = C.CURL_VERSION_IPV6
	VERSION_KERBEROS4    = C.CURL_VERSION_KERBEROS4
	VERSION_SSL          = C.CURL_VERSION_SSL
	VERSION_LIBZ         = C.CURL_VERSION_LIBZ
	VERSION_NTLM         = C.CURL_VERSION_NTLM
	VERSION_GSSNEGOTIATE = C.CURL_VERSION_GSSNEGOTIATE
	VERSION_DEBUG        = C.CURL_VERSION_DEBUG
	VERSION_ASYNCHDNS    = C.CURL_VERSION_ASYNCHDNS
	VERSION_SPNEGO       = C.CURL_VERSION_SPNEGO
	VERSION_LARGEFILE    = C.CURL_VERSION_LARGEFILE
	VERSION_IDN          = C.CURL_VERSION_IDN
	VERSION_SSPI         = C.CURL_VERSION_SSPI
	VERSION_CONV         = C.CURL_VERSION_CONV
	VERSION_CURLDEBUG    = C.CURL_VERSION_CURLDEBUG
	VERSION_TLSAUTH_SRP  = C.CURL_VERSION_TLSAUTH_SRP
	VERSION_NTLM_WB      = C.CURL_VERSION_NTLM_WB
)

// for OPT_READFUNCTION, return a int flag
const (
	READFUNC_ABORT = C.CURL_READFUNC_ABORT
	READFUNC_PAUSE = C.CURL_READFUNC_PAUSE
)

// for easy.Setopt(OPT_HTTP_VERSION, flag)
const (
	HTTP_VERSION_NONE = C.CURL_HTTP_VERSION_NONE
	HTTP_VERSION_1_0  = C.CURL_HTTP_VERSION_1_0
	HTTP_VERSION_1_1  = C.CURL_HTTP_VERSION_1_1
)

// for easy.Setopt(OPT_PROXYTYPE, flag)
const (
	PROXY_HTTP     = C.CURLPROXY_HTTP     /* added in 7.10, new in 7.19.4 default is to use CONNECT HTTP/1.1 */
	PROXY_HTTP_1_0 = C.CURLPROXY_HTTP_1_0 /* added in 7.19.4, force to use CONNECT HTTP/1.0  */
	PROXY_SOCKS4   = C.CURLPROXY_SOCKS4   /* support added in 7.15.2, enum existed already in 7.10 */
	PROXY_SOCKS5   = C.CURLPROXY_SOCKS5   /* added in 7.10 */
	PROXY_SOCKS4A  = C.CURLPROXY_SOCKS4A  /* added in 7.18.0 */
	// Use the SOCKS5 protocol but pass along the host name rather than the IP address.
	PROXY_SOCKS5_HOSTNAME = C.CURLPROXY_SOCKS5_HOSTNAME
)

// for easy.Setopt(OPT_SSLVERSION, flag)
const (
	SSLVERSION_DEFAULT = C.CURL_SSLVERSION_DEFAULT
	SSLVERSION_TLSv1   = C.CURL_SSLVERSION_TLSv1
	SSLVERSION_SSLv2   = C.CURL_SSLVERSION_SSLv2
	SSLVERSION_SSLv3   = C.CURL_SSLVERSION_SSLv3
)

// for easy.Setopt(OPT_TIMECONDITION, flag)
const (
	TIMECOND_NONE         = C.CURL_TIMECOND_NONE
	TIMECOND_IFMODSINCE   = C.CURL_TIMECOND_IFMODSINCE
	TIMECOND_IFUNMODSINCE = C.CURL_TIMECOND_IFUNMODSINCE
	TIMECOND_LASTMOD      = C.CURL_TIMECOND_LASTMOD
)

// for easy.Setopt(OPT_NETRC, flag)
const (
	NETRC_IGNORED  = C.CURL_NETRC_IGNORED
	NETRC_OPTIONAL = C.CURL_NETRC_OPTIONAL
	NETRC_REQUIRED = C.CURL_NETRC_REQUIRED
)

// for easy.Setopt(OPT_FTP_CREATE_MISSING_DIRS, flag)
const (
	FTP_CREATE_DIR_NONE  = C.CURLFTP_CREATE_DIR_NONE
	FTP_CREATE_DIR       = C.CURLFTP_CREATE_DIR
	FTP_CREATE_DIR_RETRY = C.CURLFTP_CREATE_DIR_RETRY
)

// for easy.Setopt(OPT_IPRESOLVE, flag)
const (
	IPRESOLVE_WHATEVER = C.CURL_IPRESOLVE_WHATEVER
	IPRESOLVE_V4       = C.CURL_IPRESOLVE_V4
	IPRESOLVE_V6       = C.CURL_IPRESOLVE_V6
)

// for easy.Setopt(OPT_SSL_OPTIONS, flag)
const (
	SSLOPT_ALLOW_BEAST = 1
)

// for easy.Pause(flat)
const (
	PAUSE_RECV      = C.CURLPAUSE_RECV
	PAUSE_RECV_CONT = C.CURLPAUSE_RECV_CONT
	PAUSE_SEND      = C.CURLPAUSE_SEND
	PAUSE_SEND_CONT = C.CURLPAUSE_SEND_CONT
	PAUSE_ALL       = C.CURLPAUSE_ALL
	PAUSE_CONT      = C.CURLPAUSE_CONT
)

// for multi.Info_read()
const (
	CURLMSG_NONE = C.CURLMSG_NONE
	CURLMSG_DONE = C.CURLMSG_DONE
	CURLMSG_LAST = C.CURLMSG_LAST
)

// for OPT_DEBUGFUNCTION, curl_infotype
const (
	INFOTYPE_TEXT         = C.CURLINFO_TEXT
	INFOTYPE_HEADER_IN    = C.CURLINFO_HEADER_IN
	INFOTYPE_HEADER_OUT   = C.CURLINFO_HEADER_OUT
	INFOTYPE_DATA_IN      = C.CURLINFO_DATA_IN
	INFOTYPE_DATA_OUT     = C.CURLINFO_DATA_OUT
	INFOTYPE_SSL_DATA_IN  = C.CURLINFO_SSL_DATA_IN
	INFOTYPE_SSL_DATA_OUT = C.CURLINFO_SSL_DATA_OUT
)
