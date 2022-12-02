# STIR/SHAKEN Certificate Repository Compliance

## U. S. Telepacific Corp

Name: `https://qcall.meta.tpx.net/certs/shaken_cacert.crt`\
Tested At: 02 Dec 22 07:28 UTC\
Time: 300ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 02 Dec 22 07:30 UTC