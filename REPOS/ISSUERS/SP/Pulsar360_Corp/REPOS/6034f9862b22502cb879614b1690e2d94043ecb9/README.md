# STIR/SHAKEN Certificate Repository Compliance

## Pulsar360 Corp

Name: `https://ecms.securetransit.net/certs/410bad9e05a57353f0d8739a421be619.cer`\
Tested At: 31 Oct 22 19:21 UTC\
Time: 413ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 31/10/2022 at 19:21:49