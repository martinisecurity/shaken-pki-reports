# STIR/SHAKEN Certificate Repository Compliance

## UHK Networks LLC

Name: `https://ssc.getsipnav.com/certs/35b6438acf154c4814d7c3cf8cc0f4e03d75b23b`\
Tested At: 18 Aug 25 21:05 UTC\
Time: 286ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC