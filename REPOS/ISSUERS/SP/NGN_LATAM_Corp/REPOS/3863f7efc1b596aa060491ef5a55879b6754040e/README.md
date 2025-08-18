# STIR/SHAKEN Certificate Repository Compliance

## NGN LATAM Corp

Name: `https://ssc.getsipnav.com/certs/810723d5a9f926973ae9e237f68202bf55cbb52f`\
Tested At: 18 Aug 25 21:06 UTC\
Time: 283ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC