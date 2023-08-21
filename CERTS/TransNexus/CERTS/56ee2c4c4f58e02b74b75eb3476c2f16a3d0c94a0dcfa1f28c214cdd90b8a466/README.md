# STIR/SHAKEN CA Ecosystem Compliance

## Certificate TransNexus Issuing CA G2

Tested At: 21 Aug 23 20:18 UTC\
Initial Validity Period: 4383 day(s)\
Remaining Validity Period: 3363 day(s)\
Subject: CN=TransNexus Issuing CA G2, OU=Certification Authorities, O=TransNexus, C=US\
Issuer: CN=TransNexus Root CA G2, OU=Certification Authorities, O=TransNexus, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIICSjCCAe%2BgAwIBAgIQb1EQH5xFyuPdlc4wTK2DPDAKBggqhkjOPQQDAjBmMQswCQYDVQQGEwJVUzETMBEGA1UEChMKVHJhbnNOZXh1czEiMCAGA1UECxMZQ2VydGlmaWNhdGlvbiBBdXRob3JpdGllczEeMBwGA1UEAxMVVHJhbnNOZXh1cyBSb290IENBIEcyMB4XDTIwMTEwNDAwMDAwMFoXDTMyMTEwMzIzNTk1OVowaTELMAkGA1UEBhMCVVMxEzARBgNVBAoTClRyYW5zTmV4dXMxIjAgBgNVBAsTGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxITAfBgNVBAMTGFRyYW5zTmV4dXMgSXNzdWluZyBDQSBHMjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNlfglH3GiHuzkfpNRLQFH5ptp77%2BxCnIIPczCQ%2B84Aik%2BnOeDdjFU64pehF7jJSS6kMGzvsFjbE5vZzMKMfOqqjfDB6MA4GA1UdDwEB%2FwQEAwIBBjAPBgNVHRMBAf8EBTADAQH%2FMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU%2FiVuNntTwyoXcwhCwWTGeYunJqgwHwYDVR0jBBgwFoAUd51fsT76PVNl361YAUbxtG5iZrYwCgYIKoZIzj0EAwIDSQAwRgIhAJLx79ihUjgBks75UUKVJ6Uo4KkUtWCT896FP6VRmegWAiEA9uaIZGFGabQNAVf1o6rAtN%2Fs03a6p%2BPGC62%2B0n5rUro%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ca_key_usage_crl_sign](../../ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | error | US_SHAKEN_CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |

### Not Effective

- e_atis_ca_certificate_policies
- e_atis_ca_crl_distribution
- e_atis_ca_extension_unknown
- e_atis_ca_issuer_dn
- e_atis_ca_serial_number
- e_atis_ca_subject_cn
- n_atis_ca_certificate_policy_critical

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Aug 23 20:18 UTC