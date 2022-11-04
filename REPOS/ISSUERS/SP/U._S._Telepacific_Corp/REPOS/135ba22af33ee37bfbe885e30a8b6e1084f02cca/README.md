# STIR/SHAKEN Certificate Repository Compliance

## U. S. Telepacific Corp

Name: `https://qcall.meta.tpx.net/certs/shaken_cacert.crt`\
Tested At: 04 Nov 22 01:10 UTC\
Time: 382ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Nov 22 01:11 UTC