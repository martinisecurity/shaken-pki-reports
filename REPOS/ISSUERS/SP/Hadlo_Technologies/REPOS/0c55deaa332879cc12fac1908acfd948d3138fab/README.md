# STIR/SHAKEN Certificate Repository Compliance

## Hadlo Technologies

Name: `https://www.hadlotechnologies.com/ss_certs/hadlo_ssi_public_20231113.crt`\
Tested At: 12 Feb 24 16:58 UTC\
Time: 104ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_tls_transport](../../ISSUES/e_tls_transport/README.md) | error | System | Get "https://www.hadlotechnologies.com/ss_certs/hadlo_ssi_public_20231113.crt": x509: certificate has expired or is not yet valid: current time 2024-02-12T16:58:28Z is after 2024-02-11T23:59:59Z |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 12 Feb 24 17:02 UTC