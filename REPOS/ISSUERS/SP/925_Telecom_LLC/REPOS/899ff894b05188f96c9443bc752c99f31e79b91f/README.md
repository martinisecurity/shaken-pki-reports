# STIR/SHAKEN Certificate Repository Compliance

## 925 Telecom LLC

Name: `https://ssc.getsipnav.com/certs/765ab8defea02c4a4de45048909cf0c21946154f`\
Tested At: 22 Aug 24 15:37 UTC\
Time: 59ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 22 Aug 24 15:44 UTC