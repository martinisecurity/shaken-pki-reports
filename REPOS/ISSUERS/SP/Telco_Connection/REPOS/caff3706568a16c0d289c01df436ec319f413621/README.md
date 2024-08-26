# STIR/SHAKEN Certificate Repository Compliance

## Telco Connection

Name: `https://ssc.getsipnav.com/certs/a3abaabae48b27b3d817385ff49f2858b6deaddd`\
Tested At: 26 Aug 24 18:42 UTC\
Time: 102ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 26 Aug 24 18:49 UTC