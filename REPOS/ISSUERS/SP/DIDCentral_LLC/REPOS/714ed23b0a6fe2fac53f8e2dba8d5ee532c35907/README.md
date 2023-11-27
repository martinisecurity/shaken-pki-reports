# STIR/SHAKEN Certificate Repository Compliance

## DIDCentral LLC

Name: `https://ssc.getsipnav.com/certs/92f7b56e7bc6ab17a5e4cc1d0e0e7f7eab169379`\
Tested At: 27 Nov 23 23:22 UTC\
Time: 152ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 27 Nov 23 23:28 UTC