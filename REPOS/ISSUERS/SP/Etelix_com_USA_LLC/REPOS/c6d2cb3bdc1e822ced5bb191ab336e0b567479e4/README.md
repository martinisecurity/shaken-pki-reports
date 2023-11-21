# STIR/SHAKEN Certificate Repository Compliance

## Etelix.com USA LLC

Name: `http://certs-clb-1999937273.us-east-1.elb.amazonaws.com/576K_20230815.crt`\
Tested At: 21 Nov 23 16:41 UTC\
Time: 169ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_protocol](../../ISSUES/w_atis_protocol/README.md) | warn | ATIS-1000080 | The verifier should not dereference any protocol other than https or a port other than 443 or 8443 |

Generated: 21 Nov 23 17:16 UTC