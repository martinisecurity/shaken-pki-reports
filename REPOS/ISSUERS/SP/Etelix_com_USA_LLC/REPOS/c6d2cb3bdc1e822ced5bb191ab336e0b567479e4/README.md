# STIR/SHAKEN Certificate Repository Compliance

## Etelix.com USA LLC

Name: `http://certs-clb-1999937273.us-east-1.elb.amazonaws.com/576K_20230815.crt`\
Tested At: 12 Feb 24 18:53 UTC\
Time: 23ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_protocol](../../ISSUES/w_atis_protocol/README.md) | warn | ATIS-1000080 | The verifier should not dereference any protocol other than https or a port other than 443 or 8443 |

Generated: 12 Feb 24 19:26 UTC