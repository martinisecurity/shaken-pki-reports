# STIR/SHAKEN Certificate Repository Compliance

## ANI Networks

Name: `https://certs.iverify-aninetworks.net/aninetworks_20220405.crt`\
Tested At: 05 Apr 24 18:46 UTC\
Time: 216ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 05 Apr 24 19:04 UTC