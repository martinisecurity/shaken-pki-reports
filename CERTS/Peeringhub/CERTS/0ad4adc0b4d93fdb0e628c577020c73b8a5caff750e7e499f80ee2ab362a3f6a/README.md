# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Peeringhub Inc Root CA

Tested At: 11 Jan 23 21:04 UTC\
Initial Validity Period: 7300 day(s)\
Remaining Validity Period: 6545 day(s)\
Subject: CN=Peeringhub Inc Root CA, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Issuer: CN=Peeringhub Inc Root CA, OU=Certification Authorities, O=Peeringhub Inc, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIICEDCCAbWgAwIBAgIJAJZM3gWl52yeMAoGCCqGSM49BAMCMGsxCzAJBgNVBAYTAlVTMRcwFQYDVQQKDA5QZWVyaW5naHViIEluYzEiMCAGA1UECwwZQ2VydGlmaWNhdGlvbiBBdXRob3JpdGllczEfMB0GA1UEAwwWUGVlcmluZ2h1YiBJbmMgUm9vdCBDQTAeFw0yMDEyMTcxNTMxMDRaFw00MDEyMTIxNTMxMDRaMGsxCzAJBgNVBAYTAlVTMRcwFQYDVQQKDA5QZWVyaW5naHViIEluYzEiMCAGA1UECwwZQ2VydGlmaWNhdGlvbiBBdXRob3JpdGllczEfMB0GA1UEAwwWUGVlcmluZ2h1YiBJbmMgUm9vdCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOnrCNZEAlQT3yv%2FZ2pcR%2FNtTuNf7k3pjojYhYvA83uMCzdu9KlG%2BLdDwFcioZPF027ks4CaeenplBMGhNMkVnajQjBAMA4GA1UdDwEB%2FwQEAwIBhjAPBgNVHRMBAf8EBTADAQH%2FMB0GA1UdDgQWBBQc6aJG%2F%2BsTGtSS54L0Yo7%2BBaJ%2FzzAKBggqhkjOPQQDAgNJADBGAiEA%2Frriu3nwg8fan3oUiIA%2B12PM%2Bzuj6XfE7Ay92wNItTwCIQCrkOhIJG0L2Mg%2FwpDJPTWO4gaTpGqYbdexFlsinEAaFg%3D%3D)

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


Generated: 11 Jan 23 21:04 UTC