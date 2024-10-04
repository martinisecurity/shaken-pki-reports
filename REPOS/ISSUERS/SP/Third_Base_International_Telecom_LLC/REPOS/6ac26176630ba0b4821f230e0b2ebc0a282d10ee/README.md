# STIR/SHAKEN Certificate Repository Compliance

## Third Base International Telecom LLC

Name: `https://ssc.getsipnav.com/certs/5f9a98adbc990059414e49b5b127c63af794b2a9`\
Tested At: 04 Oct 24 16:23 UTC\
Time: 40ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC