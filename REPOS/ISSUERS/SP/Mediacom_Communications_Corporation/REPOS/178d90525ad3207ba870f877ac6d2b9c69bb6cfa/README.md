# STIR/SHAKEN Certificate Repository Compliance

## Mediacom Communications Corporation

Name: `https://shaken.stir.mediacomcc.com/certs/mediacomcertchain.crt`\
Tested At: 09 Mar 23 23:18 UTC\
Time: 313ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 10 Mar 23 02:25 UTC