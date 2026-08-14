import Service from '@ember/service';
import { tracked } from '@glimmer/tracking';
import { action } from '@ember/object';

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
        },
        body: JSON.stringify({
          username_or_email: usernameOrEmail,
          password: password,
        }),
      });

      const data = await response.json();

      if (!response.ok || !data.success) {
        throw new Error(data.error || data.message || 'Login gagal. Periksa kembali kredensial Anda.');
      }

      this.token = data.token;
      this.user = data.user;

      if (typeof window !== 'undefined') {
        localStorage.setItem('aifl_admin_token', data.token);
        localStorage.setItem('aifl_admin_user', JSON.stringify(data.user));
      }

      return true;
    } catch (err) {
      this.error = err.message || 'Terjadi kesalahan saat menghubungkan ke server authentication.';
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
