# STIR/SHAKEN Certificate Repository Compliance

## IT Vocal LLC

Name: `https://shaken.vocaltransit.com/783J/vocaltransit-STIRSHAKEN.crt`\
Tested At: 01 Nov 22 20:27 UTC\
Time: 357ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 01/11/2022 at 20:34:21