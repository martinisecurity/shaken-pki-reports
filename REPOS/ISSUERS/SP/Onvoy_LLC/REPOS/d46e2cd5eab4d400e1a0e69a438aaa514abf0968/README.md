# STIR/SHAKEN Certificate Repository Compliance

## Onvoy LLC

Name: `https://sticr-cstga.ccid.neustar/api/v1/certificate/db10086649dd47f03680cc853c3a387c.pem`\
Tested At: 18 Aug 25 21:08 UTC\
Time: 111ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC