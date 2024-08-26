# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 128 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 97 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 31 certificates being tested against the remaining rules
- 2.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 29 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 31 certificates expire in the next 30 days
- 31.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 31 | [e_atis_ext_crl_distribution_struct](ISSUES/e_atis_ext_crl_distribution_struct/README.md) | ATIS1000080 |
| 31 | [e_atis_serial_number_size](ISSUES/e_atis_serial_number_size/README.md) | ATIS1000080 |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 2.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 50.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 50.00% of certificates are too old to be assessed against currently enforced expectations
- 6759 days is the average remaining validity for the certificates in the corpus
- 7300 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ext_crl_distribution_struct_ca](ISSUES/e_atis_ext_crl_distribution_struct_ca/README.md) | ATIS1000080 |
| 1 | [e_atis_serial_number_size_ca](ISSUES/e_atis_serial_number_size_ca/README.md) | ATIS1000080 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 27&#160;Jul&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 26&#160;Aug&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/0414e6f8827b9489dbf152913b716193fd91b8fa67167fd7d75f32e7c84f5f5d/README.md) |
| 28&#160;Jul&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 27&#160;Aug&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/b03d9b2da411744f4f009dceed448547f183ab184a100581edf8a4d283d2e613/README.md) |
| 29&#160;Jul&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 28&#160;Aug&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/4464420fc987e8835b52e18ff686b380ac950b2969df825bb34ac7757e4a3836/README.md) |
| 30&#160;Jul&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 29&#160;Aug&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/ceed904e5a5bae96078d27502cdb413fb8b6d3139b2ba895c7f30642f78e4d31/README.md) |
| 31&#160;Jul&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 30&#160;Aug&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/2a28e1943e87fd0beebfcfb446290fd184b955624c1dbe387e76ce842bfb12f0/README.md) |
| 01&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 31&#160;Aug&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/5e3f3ed55e427db9f40ef2a71416b3695cf5987fbfd1d4f004031c3f36c58111/README.md) |
| 02&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 01&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/1b3c95b133f54570ef88696898fce4ada0adfeb526c01bf1de4b1461bcfd6153/README.md) |
| 03&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 02&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/9c3bbc5907a7213ade2c04c274f250618955e23b1b81d6c50d980092b8e19ff5/README.md) |
| 04&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 03&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/97f0fb342bc50e7efffaf3b810d690ae99f657a101e223c704d4590293a83040/README.md) |
| 05&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 04&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/c30d01145f3f64adeac43ff7b46e8446f9907b4145b75828265a4a7467f102e9/README.md) |
| 06&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 05&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/80b23b5f69d909be3a614e98d672c099d22e7fc4579f86e6cedd622bcb8d7241/README.md) |
| 07&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 06&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/a24bac8ab5d93ac32a3be4a2c533f3967a12d6f5a93b30984bd6d2f6b57eadd5/README.md) |
| 08&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 07&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/601e0acba0436165fff10a6ad3bf529b8af7a0903364a8d86201195f4fa8201d/README.md) |
| 09&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 08&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/017566e052a421d34468b6656a81e558b4a944de27b642965075c22165a621a0/README.md) |
| 10&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 09&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/2286e840c3410161be9380f33f0ed2ea5915012ec73156e0f316da814f71d10b/README.md) |
| 11&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 10&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/4ef38c61a8a3a28749cc1e39d26e1b5eabeb5e965cef71bf675b163fc0acd522/README.md) |
| 12&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 11&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/8ad8ad8474fd189e5819cb65065243bdc4c58dcbb0e0b5a83952fe16e8923378/README.md) |
| 13&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 12&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/4e1de78b902ab9a14900ce497ee1e5bb3f814612dadac8cbfe77301ee4ddd084/README.md) |
| 14&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 13&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/88fb765cce7c5c3ac7318d902ede6c23c3e43e7c114c64b06c6f8f9619de9c9e/README.md) |
| 15&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 14&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/03c6676822190bb31410b86a936bb7cde71dd6480ababa95d47b1f0050a52b88/README.md) |
| 16&#160;Aug&#160;24&#160;06:32&#160;UTC | SHAKEN 318J | 15&#160;Sep&#160;24&#160;06:32&#160;UTC | true | [view](CERTS/2178fa58e2b64cb3e5d633e0d3b9a9e28e99aacc7b1c2dc3e2f753ca91f503d3/README.md) |
| 16&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 15&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/ac830e1de5fb354a418b99c44937cd99a4076ff86ff4e32f9afad0c0667c71a5/README.md) |
| 17&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 16&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/b14e7e74af77e292a589cdfdc5c440f679899583e10d36b009d188341467a957/README.md) |
| 18&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 17&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/ac49d78adcfb70a58492656220805f163a90695f33a73ea696bdd9567d83b841/README.md) |
| 19&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 18&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/b577618b4fb5c045f5c68d00e710177193e3118971e4447aafdd5d52b3ebf7b6/README.md) |
| 20&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 19&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/dd81139ed4fdd6f7a898727071e155b2517bacddf3a74837a74082197eab5244/README.md) |
| 21&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 20&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/a0759b28aeacb76e310496fcbd5c30593a7d1518d9cbfb46bbac606d6e810050/README.md) |
| 22&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 21&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/a7ab6b74eb35bcd8dbb877c8753ec9cbae630fa0999e010efc5d869fbe4da725/README.md) |
| 23&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 22&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/c0668fa6d4f0e3b1692f06642b2b0f0951519dbe129d39ec8ad3a7d24a8a3b9d/README.md) |
| 24&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 23&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/b9f8e17cd4daa62edfeea19a5f976f5f6bd227f364685b4a0385a2bff04a7a19/README.md) |
| 25&#160;Aug&#160;24&#160;20:55&#160;UTC | SHAKEN 318J | 24&#160;Sep&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/0c4b698bb1181519a298ba1c3ef18e2458b0f3cbd0ff02c761fd566445cfeb28/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 11&#160;Jan&#160;24&#160;02:59&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jan&#160;44&#160;02:59&#160;UTC | true | [view](CERTS/25af737667ed8b05cb8b8e7f44b2d7b5861551bea95ec48a73306ec75a92a662/README.md) |


Generated: 26 Aug 24 18:03 UTC