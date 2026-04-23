# ClientCertificate

Client TLS certificate for mTLS check endpoints.

## Details

| | |
|---|---|
| **API Group** | `infra.checkly.crossplane.io/v1alpha1` |
| **Kind** | `ClientCertificate` |
| **Terraform Resource** | [`checkly_client_certificate`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/client_certificate) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `certificateSecretRef` | `object` | (String) The client certificate in PEM format.
The client certificate in PEM format. |
| `host` | `string` | (String) The host domain that the certificate should be used for.
The host domain that the certificate should be used... |
| `passphraseSecretRef` | `object` | (String, Sensitive) Passphrase for the private key.
Passphrase for the private key. |
| `privateKeySecretRef` | `object` | (String) The private key for the certificate in PEM format.
The private key for the certificate in PEM format. |
| `trustedCa` | `string` | (String) PEM formatted bundle of CA certificates that the client should trust. The bundle may contain many CA certifi... |

## Example

```yaml
apiVersion: infra.checkly.m.crossplane.io/v1alpha1
kind: ClientCertificate
metadata:
  annotations:
    meta.upbound.io/example-id: infra/v1alpha1/clientcertificate
  labels:
    testing.upbound.io/example-name: test
  name: test
  namespace: upbound-system
spec:
  forProvider:
    certificateSecretRef:
      key: attribute.cert.pem
      name: example-secret
    host: '*.acme.com'
    passphraseSecretRef:
      key: example-key
      name: example-secret
    privateKeySecretRef:
      key: attribute.key.pem
      name: example-secret
    trustedCa: ${file("${path.module}/ca.pem")}
```

## Notes

- Both **namespace-scoped** (`infra.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`infra.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/client_certificate) for detailed field descriptions.
