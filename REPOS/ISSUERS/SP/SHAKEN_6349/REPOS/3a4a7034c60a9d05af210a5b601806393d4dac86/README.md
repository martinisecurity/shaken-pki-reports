# STIR/SHAKEN Certificate Repository Compliance

## SHAKEN 6349

Name: `https://robocall.sti.uscellular.com/certs/uscc_shaken_CA.crt`\
Tested At: 04 Oct 24 16:23 UTC\
Time: 89ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC