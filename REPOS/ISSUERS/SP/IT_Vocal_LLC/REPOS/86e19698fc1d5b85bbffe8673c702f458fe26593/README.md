# STIR/SHAKEN Certificate Repository Compliance

## IT Vocal LLC

Name: `https://shaken.vocaltransit.com/783J/vocaltransit-STIRSHAKEN.crt`\
Tested At: 10 Nov 22 23:21 UTC\
Time: 142ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 10 Nov 22 23:30 UTC