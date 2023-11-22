# STIR/SHAKEN Certificate Repository Compliance

## SipPhony LLC

Name: `https://ssc.getsipnav.com/certs/986fbcaa0958b17ff61a30c6c909f645c9ccb766`\
Tested At: 22 Nov 23 03:32 UTC\
Time: 148ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 22 Nov 23 03:38 UTC