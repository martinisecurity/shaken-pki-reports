# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 606F

Tested At: 16 Mar 23 19:03 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -117 day(s)\
Subject: CN=SHAKEN 606F, OU=SHAKEN, O=Global Data Systems Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/756cb700-f9d2-4a05-850e-c9dfe3e22de4/712fc5b5866f6ee8e9867225f645ec1a.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BTCCAqCgAwIBAgIQYSanY3ndlvV7zf3N8NxdgDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTEyMTQzMDhaFw0yMjExMTgyMTQzMDdaMFcxCzAJBgNVBAYTAlVTMSEwHwYDVQQKExhHbG9iYWwgRGF0YSBTeXN0ZW1zIEluYy4xDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDYwNkYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASfzJfpFiOj0UuZ08IRvVyxwLtpNf6WmXgaNJSXD8FYQaIxDSJRgYCYAOEgQH2pzg6D5B6QBNlfpDqVJ5SquSvFo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFC6YrRkITR7%2FdCr%2BOphZD0ilhYMCMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDYwNkYwCgYIKoZIzj0EAwIDRwAwRAIgUyz8odB1bIcXuRFw7csabja5pgY%2BpS6AnQyA1Ljp5z8CIDR5i5v%2F1GvdfxDXY%2FJDNzT0yyiFRiB48CPDwrneStoM)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 16 Mar 23 19:18 UTC