# STIR/SHAKEN Certificate Repository Compliance

## IEvolve

Name: `https://icomm.i-evolve.net/certs/icomm-shaken-sbc01.crt`\
Tested At: 01 Nov 22 09:58 UTC\
Time: 273ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 01/11/2022 at 10:05:32