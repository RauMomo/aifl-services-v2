import Service from '@ember/service';
import { tracked } from '@glimmer/tracking';
import { action } from '@ember/object';
import protobuf from 'protobufjs';

const adminProtoSchema = {
  nested: {
    pb: {
      nested: {
        AdminUserProto: {
          fields: {
            id: { type: 'string', id: 1 },
            username: { type: 'string', id: 2 },
            email: { type: 'string', id: 3 },
            role: { type: 'string', id: 4 },
            created_at: { type: 'string', id: 5 },
          },
        },
        AdminLoginResponseProto: {
          fields: {
            success: { type: 'bool', id: 1 },
            message: { type: 'string', id: 2 },
            token: { type: 'string', id: 3 },
            user: { type: 'AdminUserProto', id: 4 },
            error: { type: 'string', id: 5 },
          },
        },
      },
    },
  },
};

const root = protobuf.Root.fromJSON(adminProtoSchema);
const AdminLoginResponseProto = root.lookupType('pb.AdminLoginResponseProto');

export default class AuthService extends Service {
  @tracked token = null;
  @tracked user = null;
  @tracked isLoading = false;
  @tracked error = null;

  constructor() {
    super(...arguments);
    if (typeof window !== 'undefined') {
      const savedToken = localStorage.getItem('aifl_admin_token');
      const savedUser = localStorage.getItem('aifl_admin_user');
      if (savedToken) {
        this.token = savedToken;
        if (savedUser) {
          try {
            this.user = JSON.parse(savedUser);
          } catch {
            this.user = null;
          }
        }
      }
    }
  }

  get isAuthenticated() {
    return Boolean(this.token);
  }

  @action
  async login(usernameOrEmail, password) {
    this.isLoading = true;
    this.error = null;

    try {
      const apiHost =
        import.meta.env.VITE_BACKEND_URL || 'http://localhost:8200';
      const response = await fetch(`${apiHost}/api/v1/auth/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/x-protobuf',
        },
        body: JSON.stringify({
          username_or_email: usernameOrEmail,
          password: password,
        }),
      });

      let data;
      const contentType = response.headers.get('content-type');
      if (contentType && contentType.includes('application/x-protobuf')) {
        const buffer = await response.arrayBuffer();
        const decoded = AdminLoginResponseProto.decode(new Uint8Array(buffer));
        data = AdminLoginResponseProto.toObject(decoded, { defaults: true });
      } else {
        data = await response.json();
      }

      if (!response.ok || !data.success) {
        throw new Error(
          data.error ||
            data.message ||
            'Login gagal. Periksa kembali kredensial Anda.',
        );
      }

      this.token = data.token;
      this.user = data.user;

      if (typeof window !== 'undefined') {
        localStorage.setItem('aifl_admin_token', data.token);
        localStorage.setItem('aifl_admin_user', JSON.stringify(data.user));
      }

      return true;
    } catch (err) {
      this.error =
        err.message ||
        'Terjadi kesalahan saat menghubungkan ke server authentication.';
      return false;
    } finally {
      this.isLoading = false;
    }
  }

  @action
  logout() {
    this.token = null;
    this.user = null;
    if (typeof window !== 'undefined') {
      localStorage.removeItem('aifl_admin_token');
      localStorage.removeItem('aifl_admin_user');
    }
  }
}
