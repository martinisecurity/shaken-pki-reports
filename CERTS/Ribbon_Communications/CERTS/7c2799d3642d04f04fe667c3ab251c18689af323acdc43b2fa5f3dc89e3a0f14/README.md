# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Ribbon Root CA

Tested At: 23 Nov 22 18:09 UTC\
Initial Validity Period: 9131 day(s)\
Remaining Validity Period: 8572 day(s)\
Subject: CN=SHAKEN Ribbon Root CA, OU=Certification Authority, O=Ribbon Communications, C=US\
Issuer: CN=SHAKEN Ribbon Root CA, OU=Certification Authority, O=Ribbon Communications, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIICHzCCAcSgAwIBAgIQdbMkHP%2FQn78dmgdxeLv0djAKBggqhkjOPQQDAjBvMQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSAwHgYDVQQLExdDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eTEeMBwGA1UEAxMVU0hBS0VOIFJpYmJvbiBSb290IENBMB4XDTIxMDUxMzAwMDAwMFoXDTQ2MDUxMjIzNTk1OVowbzELMAkGA1UEBhMCVVMxHjAcBgNVBAoTFVJpYmJvbiBDb21tdW5pY2F0aW9uczEgMB4GA1UECxMXQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkxHjAcBgNVBAMTFVNIQUtFTiBSaWJib24gUm9vdCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCY6GnVTWMs4FS83ne8vp9hz6F1Vk%2BthRDGpSb1jaLa8ZWhUtVbcUfoE%2F1bw3kspDyjuG0%2BjkJUaKqEmQkYnDOSjQjBAMA8GA1UdEwEB%2FwQFMAMBAf8wDgYDVR0PAQH%2FBAQDAgEGMB0GA1UdDgQWBBS7bhaSFwYNkmN%2BwZQyNeVu0dUUZDAKBggqhkjOPQQDAgNJADBGAiEAucCJi5Vvv0xHWAr6GXzyiZqG%2FbiL%2Bcrcr1JHVCmMbxICIQC8wkpOWO4DJpA3N54ig2xCp5dsoiibF4BJx64gYCZLww%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ca_key_usage_crl_sign](../../ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | error | US_SHAKEN_CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |

### Not Effective

- e_atis_ca_serial_number
- e_atis_ca_subject_cn
- e_atis_root_certificate_policies
- e_atis_root_extension_unknown

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 23 Nov 22 18:09 UTC