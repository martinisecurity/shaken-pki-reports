# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate ba68a7b635a7e85fc1ff99d8f65d7f0ea640dcbd
Tested At: 2022-10-28 19:22:02 +0000 UTC\
Initial Validity Period: 7305 day(s)\
Remaining Validity Period: 6555 day(s)\
Subject: CN=Neustar Canada Certified Caller ID SHAKEN Root CA, OU=cms-ca.ccid.neustar, O=Neustar Information Services Inc, C=CA\
Issuer: CN=Neustar Canada Certified Caller ID SHAKEN Root CA, OU=cms-ca.ccid.neustar, O=Neustar Information Services Inc, C=CA

Link: https://stir.fibernetics.ca/prod-cert2022.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIICmzCCAkKgAwIBAgIUQVJ25K%2FBQU4Tby6TTgieMvfdS6owCgYIKoZIzj0EAwIwgZIxCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEcMBoGA1UECwwTY21zLWNhLmNjaWQubmV1c3RhcjE6MDgGA1UEAwwxTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBJRCBTSEFLRU4gUm9vdCBDQTAeFw0yMDEwMDgxNDQ5NDlaFw00MDEwMDgxNDQ5NDlaMIGSMQswCQYDVQQGEwJDQTEpMCcGA1UECgwgTmV1c3RhciBJbmZvcm1hdGlvbiBTZXJ2aWNlcyBJbmMxHDAaBgNVBAsME2Ntcy1jYS5jY2lkLm5ldXN0YXIxOjA4BgNVBAMMMU5ldXN0YXIgQ2FuYWRhIENlcnRpZmllZCBDYWxsZXIgSUQgU0hBS0VOIFJvb3QgQ0EwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASjwCO8A6fzPwY1F5MuaXLpJ3idemAo1q8CdAyQXS6Ln4yVkNhghSM%2BEhKjc7RlBX340zfExlIaIzpyHTZVNvpUo3QwcjAPBgNVHRMBAf8EBTADAQH%2FMB8GA1UdIwQYMBaAFEgDf%2B8c7YwuP4vEBHIWEAP62tkuMA8GA1UdJQQIMAYGBFUdJQAwHQYDVR0OBBYEFEgDf%2B8c7YwuP4vEBHIWEAP62tkuMA4GA1UdDwEB%2FwQEAwIBhjAKBggqhkjOPQQDAgNHADBEAiARILOdPc23%2BPI4QwMZXudEjoNvusfK%2B5RMKwDnQu5LMAIgf5PAGmqtyAdFlRVuJ00qWw7G2%2BHw5Oq3NuHYjFZfzKE%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_pki_ca_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |

### Not Effective

- e_sti_ca_issuer_dn
- e_sti_ca_subject_cn
- e_sti_ca_serial_number
- e_sti_root_certificate_policies
- e_cp1_3_ca_key_usage_crl_sign
- e_sti_root_extension_unknown

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 28/10/2022 at 19:22:10