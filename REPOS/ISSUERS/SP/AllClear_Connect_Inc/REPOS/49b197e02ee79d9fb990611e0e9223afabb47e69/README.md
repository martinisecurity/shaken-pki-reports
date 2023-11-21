# STIR/SHAKEN Certificate Repository Compliance

## AllClear Connect Inc

Name: `https://ssc.getsipnav.com/certs/b4ec00e21e2d97c0707f1c726d5358dd1ae3903b`\
Tested At: 21 Nov 23 01:49 UTC\
Time: 169ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 21 Nov 23 01:55 UTC