import Component from '@glimmer/component';
import { tracked } from '@glimmer/tracking';
import { action } from '@ember/object';
import { service } from '@ember/service';
import { on } from '@ember/modifier';
import UiBadge from './ui-badge';
import UiButton from './ui-button';

export default class AdminDashboardComponent extends Component {
  @service auth;
  @service router;

  @tracked contacts = [];
  @tracked isLoadingContacts = false;
  @tracked contactError = null;

  constructor() {
    super(...arguments);
    this.fetchContacts();
  }

  @action
  async fetchContacts() {
    if (!this.auth.token) return;
    this.isLoadingContacts = true;
    this.contactError = null;

    try {
      const apiHost =
        import.meta.env.VITE_BACKEND_URL || 'http://localhost:8200';
      const res = await fetch(`${apiHost}/api/v1/contact/list`, {
        headers: {
          Authorization: `Bearer ${this.auth.token}`,
        },
      });

      const data = await res.json();
      if (data.success && data.data) {
        this.contacts = data.data;
      } else {
        this.contactError = data.error || 'Gagal mengambil data kontak.';
      }
    } catch (err) {
      this.contactError = 'Gagal terhubung ke backend server API.';
    } finally {
      this.isLoadingContacts = false;
    }
  }

  @action
  handleLogout() {
    this.auth.logout();
    this.router.transitionTo('portal-access.index');
  }

  <template>
    <div class="admin-dashboard-container fade-in">
      {{! Top Control Bar }}
      <header class="admin-top-bar">
        <div class="admin-logo-group">
          <UiBadge @variant="fill">AIFL CONTROL PORTAL</UiBadge>
          <span class="admin-bar-title">Enterprise Admin Control Center</span>
        </div>

        <div class="admin-user-group">
          <span class="admin-user-name">👤 {{this.auth.user.username}}</span>
          <span class="admin-user-role">({{this.auth.user.role}})</span>
          <button type="button" class="admin-logout-btn" {{on "click" this.handleLogout}}>
            Keluar (Logout) ➔
          </button>
        </div>
      </header>

      <main class="admin-dashboard-content">
        {{! Stats Overview Cards }}
        <div class="admin-stats-grid">
          <div class="admin-stat-card">
            <div class="stat-label">TOTAL PESAN MASUK</div>
            <div class="stat-value">{{this.contacts.length}}</div>
            <div class="stat-meta">Konsultasi Terdaftar</div>
          </div>
          <div class="admin-stat-card">
            <div class="stat-label">DATABASE ENCRYPTED</div>
            <div class="stat-value green">GORM SQLite</div>
            <div class="stat-meta">Protected Schema</div>
          </div>
          <div class="admin-stat-card">
            <div class="stat-label">CACHE & RATE LIMIT</div>
            <div class="stat-value cyan">Redis In-Memory</div>
            <div class="stat-meta">Brute-Force Protected</div>
          </div>
          <div class="admin-stat-card">
            <div class="stat-label">SERIALIZATION</div>
            <div class="stat-value amber">JSON & Protobuf</div>
            <div class="stat-meta">High Performance</div>
          </div>
        </div>

        {{! Contact Messages Table }}
        <div class="admin-panel-card">
          <div class="panel-header">
            <h3>📩 Laporan & Pesan Konsultasi Masuk</h3>
            <UiButton @variant="fill" {{on "click" this.fetchContacts}}>
              Refresh Data
            </UiButton>
          </div>

          {{#if this.isLoadingContacts}}
            <div class="admin-loading-state">
              <div class="spinner"></div>
              <span>Memuat data pesan dari backend database...</span>
            </div>
          {{else if this.contactError}}
            <div class="admin-error-box">
              <span>⚠️ {{this.contactError}}</span>
            </div>
          {{else if (eq this.contacts.length 0)}}
            <div class="admin-empty-state">
              <span>Belum ada pesan konsultasi masuk saat ini.</span>
            </div>
          {{else}}
            <div class="table-responsive">
              <table class="admin-table">
                <thead>
                  <tr>
                    <th>Waktu</th>
                    <th>Nama</th>
                    <th>Email</th>
                    <th>Perusahaan</th>
                    <th>Subjek / Pesan</th>
                    <th>Status</th>
                  </tr>
                </thead>
                <tbody>
                  {{#each this.contacts as |item|}}
                    <tr>
                      <td class="cell-time">{{item.created_at}}</td>
                      <td class="cell-name">{{item.name}}</td>
                      <td class="cell-email">{{item.email}}</td>
                      <td>{{if item.company item.company "-"}}</td>
                      <td class="cell-message">
                        <strong>{{item.subject}}</strong>
                        <p>{{item.message}}</p>
                      </td>
                      <td>
                        <span class="status-badge unread">{{item.status}}</span>
                      </td>
                    </tr>
                  {{/each}}
                </tbody>
              </table>
            </div>
          {{/if}}
        </div>
      </main>
    </div>
  </template>
}
