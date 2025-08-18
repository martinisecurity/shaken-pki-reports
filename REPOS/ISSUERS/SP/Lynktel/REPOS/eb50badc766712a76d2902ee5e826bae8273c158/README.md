# STIR/SHAKEN Certificate Repository Compliance

## Lynktel

Name: `https://ssc.getsipnav.com/certs/adb11abc916d55f59390aebfac21a502f9e4a3ef`\
Tested At: 18 Aug 25 21:06 UTC\
Time: 605ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC