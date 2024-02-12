# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Nuvera1442

Tested At: 12 Feb 24 16:32 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 47 day(s)\
Subject: C=US, ST=MN, L=Ulm, O=Nuvera, OU=Nuvera, CN=Nuvera1442\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/27.197

[View certificate details](https://understandingwebpki.com/?cert=MIID6jCCAtKgAwIBAgIUbaiefFPtePRnyVT4G94fhxLlvAcwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAzMzAxNDQwMjhaFw0yNDAzMzAxNDQwMjhaMF8xEzARBgNVBAMMCk51dmVyYTE0NDIxDzANBgNVBAsMBk51dmVyYTEPMA0GA1UECgwGTnV2ZXJhMQwwCgYDVQQHDANVbG0xCzAJBgNVBAgMAk1OMQswCQYDVQQGEwJVUzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAQnacdBhoxgaHh9rEp4z%2Fs2Gscxw86z733PnU7DsTTAKRxcUZQkUG98jN4BJ4ABbFF8cF0vN2UXBWAxTeYuWEajggFIMIIBRDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFDu5XMsxF5xIz9TaZuuNtqjV4Qn3MIGBBggrBgEFBQcBAQR1MHMwRwYIKwYBBQUHMAKGO2h0dHA6Ly9jYWNlcnRzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0MCgGCCsGAQUFBzABhhxodHRwOi8vb2NzcC1jYTEuY2NpZC5uZXVzdGFyMEgGA1UdHwRBMD8wPaA7oDmGN2h0dHA6Ly9jcmwuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcmwwHQYDVR0OBBYEFKk8nXo5v86p1MRgP89iFUle6BTxMA4GA1UdDwEB%2FwQEAwIHgDAWBggrBgEFBQcBGgQKMAigBhYEMTQ0MjANBgkqhkiG9w0BAQsFAAOCAQEAl6%2BmKSFEqUZzTnBBTGXOgMY%2B9sK4%2FEDRxZJod95PONr9lhHNHBqrRsdCpAVb3embYTWq8DVCJtpkH2wW04QF7CaYOyYeatQhdw1AT%2F0nOcmSdzpOLYSwZnRLL3iAEZf7ecu60wNBub7O3QZMV2elr6RRoiB%2BcHC4zqhGPfpndQSqjN3k1yYREKOST8CcBAhN3Ho%2BhdjYTiAKdJQQQp0eV7RyFgo8f5cU9ubcGIpRleXrSvmsF82zYFotT0rvKtU9AKcu9tiO3fUYOoKd1atTljdPR9TVEN8XMe2AWLGX%2BpcXZGjD9xHRDmGWnCBLY7XQEhz0YfqiXrEUKLWQVEbcKA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'Nuvera1442' does not contain 'SHAKEN' |
| [e_atis_signature_algorithm](../../ISSUES/e_atis_signature_algorithm/README.md) | error | ATIS1000080 | SignatureAlgorithm field is not 'ecdsa-with-SHA256', got 1.2.840.113549.1.1.11 |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Feb 24 17:02 UTC