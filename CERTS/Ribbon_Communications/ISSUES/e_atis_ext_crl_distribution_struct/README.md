# STIR/SHAKEN CA Ecosystem Compliance

## Ribbon Communications

Name: e_atis_ext_crl_distribution_struct\
Source: ATIS1000080\
Citation: ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements\
Effective Date: 16 Jan 22 00:00 UTC\
Description: STI intermediate and end-entity certificates shall contain a CRL Distribution Points extension containing a single DistributionPoint entry. The DistributionPoint entry shall contain a distributionPoint field identifying the HTTP URL reference to the file containing the SHAKEN CRL hosted by the STI-PA, and a cRLIssuer field that contains the DN of the issuer of the CRL.

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | ColoradoValley SHAKEN 2059 | [view](../../CERTS/328011b837c2166f113664ab01dc719816e4bfcb734782636e9abb7bbf68199a/README.md) | CRL Distribution Point shall contain a CRLIssuer field |
| error | Localtel SHAKEN 3229 | [view](../../CERTS/92063be4763362c7e542a41e3f4ebd030f1816f5ec8600b444be908c646307e5/README.md) | CRL Distribution Point shall contain a CRLIssuer field |
| error | Ribbon SHAKEN 2086 | [view](../../CERTS/a8850907d694bd9b9787d39dd18aff77488f3ff1d12a045b5a112dde73d8d5b2/README.md) | CRL Distribution Point shall contain a CRLIssuer field |
| error | Peoples SHAKEN 2130 | [view](../../CERTS/506e5980d9b8f24ac2c781f15c58fcc50dd50c9f82d8aac2ac72967211f262db/README.md) | CRL Distribution Point shall contain a CRLIssuer field |
| error | FirstDigital SHAKEN 5727 | [view](../../CERTS/cf281053153a12955be41e5b0f99caf1572a1c30b21f62a5c22543bf337740f7/README.md) | CRL Distribution Point shall contain a CRLIssuer field |
| error | Ironton SHAKEN 1234 0175 | [view](../../CERTS/897d53482f78c037e9dd54cc3e62a1a3dba25a5abff81f1faf3ec09ef58bb5a0/README.md) | CRL Distribution Point shall contain a CRLIssuer field |

### CA Certificates

No error, warning, or notice level issues were found


Generated: 22 Aug 24 15:44 UTC