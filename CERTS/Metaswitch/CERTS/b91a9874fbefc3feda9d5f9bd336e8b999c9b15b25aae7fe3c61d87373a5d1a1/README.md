# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Metaswitch STI-CA SHAKEN Issuing 1

Tested At: 27 Nov 23 22:55 UTC\
Initial Validity Period: 4380 day(s)\
Remaining Validity Period: 3283 day(s)\
Subject: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Issuer: CN=Metaswitch STI-CA SHAKEN Root

[View certificate details](https://understandingwebpki.com/?cert=MIIBtjCCAVygAwIBAgIQaQRv6d8JnAo38%2Ffj5Kf%2F9zAKBggqhkjOPQQDAjAoMSYwJAYDVQQDDB1NZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gUm9vdDAeFw0yMDExMjUxMTU3MzhaFw0zMjExMjIxMTU3MzhaMC0xKzApBgNVBAMMIk1ldGFzd2l0Y2ggU1RJLUNBIFNIQUtFTiBJc3N1aW5nIDEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT8rXFLu6%2B6yRdl8bcBkaYj8zn05JVmeRYW1wPXttNOhQWK%2BbYC0lE5tR3IpgMou0xcLoubgVxr4kIsVo1UmMNZo2MwYTAdBgNVHQ4EFgQUzR6nABAQ2jIdaRo51dJGCyw8h9YwHwYDVR0jBBgwFoAUOVRpugxs9p9X0ZsBMUWKoPv%2FeGowDwYDVR0TAQH%2FBAUwAwEB%2FzAOBgNVHQ8BAf8EBAMCAgQwCgYIKoZIzj0EAwIDSAAwRQIgfEI4047AI8guEw0olNkV5AFj4%2B1vUg5MbvMcQmwhu5ICIQCkBy07hzIG%2BnXxe6lxQdEKlvAEOupeV%2BGBta%2BvvtKgHw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_ca](../../ISSUES/e_atis_ext_crl_distribution_ca/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_subject_dn_ca](../../ISSUES/e_atis_subject_dn_ca/README.md) | error | ATIS1000080 | Subject DN does not contain a Country (C=) attribute |
| [e_atis_ext_certificate_policies_ca](../../ISSUES/e_atis_ext_certificate_policies_ca/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |

### Not Effective

- e_atis_ext_not_specified_ca
- e_atis_serial_number_size_ca
- e_atis_subject_c_iso_ca
- e_atis_subject_key_identifier_size_ca
- e_atis_subject_o_required_ca
- e_shaken_certificate_policies_id_ca

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 27 Nov 23 22:56 UTC