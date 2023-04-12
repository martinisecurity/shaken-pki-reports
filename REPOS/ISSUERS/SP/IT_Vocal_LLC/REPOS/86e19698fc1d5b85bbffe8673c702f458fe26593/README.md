# STIR/SHAKEN Certificate Repository Compliance

## IT Vocal LLC

Name: `https://shaken.vocaltransit.com/783J/vocaltransit-STIRSHAKEN.crt`\
Tested At: 12 Apr 23 21:52 UTC\
Time: 240ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 12 Apr 23 22:02 UTC