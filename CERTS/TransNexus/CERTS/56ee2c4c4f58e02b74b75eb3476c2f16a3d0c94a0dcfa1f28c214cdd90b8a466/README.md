# STIR/SHAKEN CA Ecosystem Compliance

## Certificate TransNexus Issuing CA G2

Tested At: 12 Feb 24 17:00 UTC\
Initial Validity Period: 4383 day(s)\
Remaining Validity Period: 3188 day(s)\
Subject: CN=TransNexus Issuing CA G2, OU=Certification Authorities, O=TransNexus, C=US\
Issuer: CN=TransNexus Root CA G2, OU=Certification Authorities, O=TransNexus, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIICSjCCAe%2BgAwIBAgIQb1EQH5xFyuPdlc4wTK2DPDAKBggqhkjOPQQDAjBmMQswCQYDVQQGEwJVUzETMBEGA1UEChMKVHJhbnNOZXh1czEiMCAGA1UECxMZQ2VydGlmaWNhdGlvbiBBdXRob3JpdGllczEeMBwGA1UEAxMVVHJhbnNOZXh1cyBSb290IENBIEcyMB4XDTIwMTEwNDAwMDAwMFoXDTMyMTEwMzIzNTk1OVowaTELMAkGA1UEBhMCVVMxEzARBgNVBAoTClRyYW5zTmV4dXMxIjAgBgNVBAsTGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxITAfBgNVBAMTGFRyYW5zTmV4dXMgSXNzdWluZyBDQSBHMjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNlfglH3GiHuzkfpNRLQFH5ptp77%2BxCnIIPczCQ%2B84Aik%2BnOeDdjFU64pehF7jJSS6kMGzvsFjbE5vZzMKMfOqqjfDB6MA4GA1UdDwEB%2FwQEAwIBBjAPBgNVHRMBAf8EBTADAQH%2FMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU%2FiVuNntTwyoXcwhCwWTGeYunJqgwHwYDVR0jBBgwFoAUd51fsT76PVNl361YAUbxtG5iZrYwCgYIKoZIzj0EAwIDSQAwRgIhAJLx79ihUjgBks75UUKVJ6Uo4KkUtWCT896FP6VRmegWAiEA9uaIZGFGabQNAVf1o6rAtN%2Fs03a6p%2BPGC62%2B0n5rUro%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_ca](../../ISSUES/e_atis_subject_cn_ca/README.md) | error | ATIS1000080 | Common Name attribute 'TransNexus Issuing CA G2' does not contain 'SHAKEN' |
| [e_atis_ext_crl_distribution_ca](../../ISSUES/e_atis_ext_crl_distribution_ca/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |

### Not Effective

- e_atis_ext_not_specified_ca
- e_atis_serial_number_size_ca
- e_atis_subject_c_iso_ca
- e_atis_subject_key_identifier_size_ca
- e_atis_subject_o_required_ca
- e_shaken_certificate_policies_id_ca

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Feb 24 17:02 UTC