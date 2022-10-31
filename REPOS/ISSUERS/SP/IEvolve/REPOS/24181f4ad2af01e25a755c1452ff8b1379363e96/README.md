# STIR/SHAKEN Certificate Repository Compliance

## IEvolve

Name: `https://icomm.i-evolve.net/certs/icomm-shaken-sbc01.crt`\
Tested At: 31 Oct 22 16:42 UTC\
Time: 423ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 31/10/2022 at 16:43:22