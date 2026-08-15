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

  @tracked activeTab = 'contacts'; // Default active tab is Contacts

  @tracked contacts = [];
  @tracked isLoadingContacts = false;
  @tracked contactError = null;

  @tracked projects = [];
  @tracked isLoadingProjects = false;
  @tracked projectError = null;

  @tracked testimonials = [];
  @tracked isLoadingTestimonials = false;
  @tracked testimonialError = null;

  teamMembers = [
    {
      name: 'Dr. Hendry Prasetya, M.T.',
      role: 'Chief Technology Officer & Lead System Architect',
      email: 'hendry@aiflservices.com',
      avatar: '👨‍💻',
      specialty: 'High-Throughput Distributed Systems & Golang Architecture',
    },
    {
      name: 'Rina Wijaya, S.Kom.',
      role: 'Head of Product & Workflow Engineering',
      email: 'rina@aiflservices.com',
      avatar: '👩‍💻',
      specialty: 'Digital Process Discovery & Ember.js Ecosystem',
    },
    {
      name: 'Budi Santoso, M.Sc.',
      role: 'Lead Cloud Infrastructure & Security Auditor',
      email: 'budi@aiflservices.com',
      avatar: '👨‍🔧',
      specialty: 'Redis Caching, Docker Orchestration & Protobuf Protocol',
    },
  ];

  constructor() {
    super(...arguments);
    this.fetchContacts();
  }

  get isContactsTab() {
    return this.activeTab === 'contacts';
  }

  get isProjectsTab() {
    return this.activeTab === 'projects';
  }

  get isTestimonialsTab() {
    return this.activeTab === 'testimonials';
  }

  get isTeamTab() {
    return this.activeTab === 'team';
  }

  @action
  selectContacts() {
    this.activeTab = 'contacts';
    this.fetchContacts();
  }

  @action
  selectProjects() {
    this.activeTab = 'projects';
    this.fetchProjects();
  }

  @action
  selectTestimonials() {
    this.activeTab = 'testimonials';
    this.fetchTestimonials();
  }

  @action
  selectTeam() {
    this.activeTab = 'team';
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
    } catch {
      this.contactError = 'Gagal terhubung ke backend server API.';
    } finally {
      this.isLoadingContacts = false;
    }
  }

  @action
  async fetchProjects() {
    this.isLoadingProjects = true;
    this.projectError = null;

    try {
      const apiHost =
        import.meta.env.VITE_BACKEND_URL || 'http://localhost:8200';
      const res = await fetch(`${apiHost}/api/v1/projects`);
      const data = await res.json();
      if (data.success && data.data) {
        this.projects = data.data;
      } else {
        this.projectError = data.error || 'Gagal mengambil data proyek.';
      }
    } catch {
      this.projectError = 'Gagal terhubung ke backend server API.';
    } finally {
      this.isLoadingProjects = false;
    }
  }

  @action
  async fetchTestimonials() {
    this.isLoadingTestimonials = true;
    this.testimonialError = null;

    try {
      const apiHost =
        import.meta.env.VITE_BACKEND_URL || 'http://localhost:8200';
      const res = await fetch(`${apiHost}/api/v1/testimonials`);
      const data = await res.json();
      if (data.success && data.data) {
        this.testimonials = data.data;
      } else {
        this.testimonialError = data.error || 'Gagal mengambil data testimoni.';
      }
    } catch {
      this.testimonialError = 'Gagal terhubung ke backend server API.';
    } finally {
      this.isLoadingTestimonials = false;
    }
  }

  @action
  handleLogout() {
    this.auth.logout();
    this.router.transitionTo('portal-access.index');
  }

  <template>
    <div class="admin-dashboard-container fade-in">
      {{! Header Bar with Left-Aligned Menus (No Sidebar) }}
      <header class="admin-top-bar">
        <div class="admin-left-header-group">
          <div class="admin-brand-header">
            <img src="/logo.png" alt="AIFL Services" class="admin-brand-logo-img" />
            <div class="admin-brand-text">
              <span class="admin-brand-name">AIFL <span class="brand-name-highlight">SERVICES</span></span>
              <span class="admin-brand-tag">Admin Portal</span>
            </div>
          </div>

          {{! Left Aligned Header Navigation Menu }}
          <nav class="admin-header-nav">
            <button
              type="button"
              class="admin-nav-btn {{if this.isContactsTab 'active'}}"
              {{on "click" this.selectContacts}}
            >
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M22 17a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8z"></path>
                <polyline points="22,9 12,15 2,9"></polyline>
              </svg>
              <span>Contacts</span>
            </button>

            <button
              type="button"
              class="admin-nav-btn {{if this.isProjectsTab 'active'}}"
              {{on "click" this.selectProjects}}
            >
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path>
              </svg>
              <span>Projects</span>
            </button>

            <button
              type="button"
              class="admin-nav-btn {{if this.isTestimonialsTab 'active'}}"
              {{on "click" this.selectTestimonials}}
            >
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
                <polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"></polygon>
              </svg>
              <span>Testimonials</span>
            </button>

            <button
              type="button"
              class="admin-nav-btn {{if this.isTeamTab 'active'}}"
              {{on "click" this.selectTeam}}
            >
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                <circle cx="9" cy="7" r="4"></circle>
                <path d="M23 21v-2a4 4 0 0 0-3-3.87"></path>
                <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
              </svg>
              <span>Our Team</span>
            </button>
          </nav>
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
        {{#if this.isContactsTab}}
          {{! 1. CONTACTS PANEL (Initial Active Page) }}
          <div class="admin-panel-card">
            <div class="panel-header">
              <div>
                <h3>📩 Pesan & Konsultasi Masuk (Contact Us)</h3>
                <p class="panel-subtitle">Daftar laporan permohonan konsultasi dari formulir publik</p>
              </div>
              <UiButton @variant="fill" {{on "click" this.fetchContacts}}>
                Refresh Data
              </UiButton>
            </div>

            {{#if this.isLoadingContacts}}
              <div class="admin-loading-state">
                <div class="spinner"></div>
                <span>Memuat data pesan dari database...</span>
              </div>
            {{else if this.contactError}}
              <div class="admin-error-box">
                <span>⚠️ {{this.contactError}}</span>
              </div>
            {{else if this.contacts.length}}
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
            {{else}}
              <div class="admin-empty-state">
                <span>Belum ada pesan konsultasi masuk saat ini.</span>
              </div>
            {{/if}}
          </div>
        {{else if this.isProjectsTab}}
          {{! 2. PROJECTS PANEL }}
          <div class="admin-panel-card">
            <div class="panel-header">
              <div>
                <h3>🚀 Daftar Proyek & Studi Kasus</h3>
                <p class="panel-subtitle">Katalog proyek solusi digital yang telah dikerjakan oleh AIFL</p>
              </div>
              <UiButton @variant="fill" {{on "click" this.fetchProjects}}>
                Refresh Proyek
              </UiButton>
            </div>

            {{#if this.isLoadingProjects}}
              <div class="admin-loading-state">
                <div class="spinner"></div>
                <span>Memuat data proyek dari backend database...</span>
              </div>
            {{else if this.projectError}}
              <div class="admin-error-box">
                <span>⚠️ {{this.projectError}}</span>
              </div>
            {{else if this.projects.length}}
              <div class="table-responsive">
                <table class="admin-table">
                  <thead>
                    <tr>
                      <th>Proyek & Kategori</th>
                      <th>Klien</th>
                      <th>Tags Teknologis</th>
                      <th>Dampak / Hasil</th>
                      <th>Unggulan</th>
                    </tr>
                  </thead>
                  <tbody>
                    {{#each this.projects as |p|}}
                      <tr>
                        <td class="cell-name">
                          <strong>{{p.title}}</strong>
                          <span class="cell-sub-category">{{p.category}}</span>
                        </td>
                        <td>{{p.client_name}}</td>
                        <td>
                          <div class="tags-container">
                            {{#each p.tags as |tag|}}
                              <span class="admin-tag-pill">{{tag}}</span>
                            {{/each}}
                          </div>
                        </td>
                        <td>
                          <div class="impact-container">
                            {{#each p.impact_metrics as |metric|}}
                              <span class="admin-metric-badge">{{metric}}</span>
                            {{/each}}
                          </div>
                        </td>
                        <td>
                          {{#if p.is_featured}}
                            <span class="status-badge featured">⭐ Featured</span>
                          {{else}}
                            <span class="status-badge standard">Standard</span>
                          {{/if}}
                        </td>
                      </tr>
                    {{/each}}
                  </tbody>
                </table>
              </div>
            {{else}}
              <div class="admin-empty-state">
                <span>Belum ada data proyek terdaftar.</span>
              </div>
            {{/if}}
          </div>
        {{else if this.isTestimonialsTab}}
          {{! 3. TESTIMONIALS PANEL }}
          <div class="admin-panel-card">
            <div class="panel-header">
              <div>
                <h3>⭐ Testimoni Klien & Ulasan Ops</h3>
                <p class="panel-subtitle">Kutipan dan ulasan pengalaman kerja sama dari klien enterprise</p>
              </div>
              <UiButton @variant="fill" {{on "click" this.fetchTestimonials}}>
                Refresh Testimoni
              </UiButton>
            </div>

            {{#if this.isLoadingTestimonials}}
              <div class="admin-loading-state">
                <div class="spinner"></div>
                <span>Memuat data testimoni dari backend database...</span>
              </div>
            {{else if this.testimonialError}}
              <div class="admin-error-box">
                <span>⚠️ {{this.testimonialError}}</span>
              </div>
            {{else if this.testimonials.length}}
              <div class="admin-cards-grid">
                {{#each this.testimonials as |t|}}
                  <div class="admin-testimonial-card">
                    <div class="testi-rating">⭐⭐⭐⭐⭐ ({{t.rating}}/5)</div>
                    <blockquote class="testi-quote">“{{t.quote}}”</blockquote>
                    <div class="testi-author-info">
                      <strong class="author-name">{{t.author_name}}</strong>
                      <span class="author-role">{{t.position}} - {{t.company}}</span>
                    </div>
                  </div>
                {{/each}}
              </div>
            {{else}}
              <div class="admin-empty-state">
                <span>Belum ada testimoni klien.</span>
              </div>
            {{/if}}
          </div>
        {{else if this.isTeamTab}}
          {{! 4. OUR TEAM PANEL }}
          <div class="admin-panel-card">
            <div class="panel-header">
              <div>
                <h3>👥 Tim Manajemen & Arsitek Operasional</h3>
                <p class="panel-subtitle">Personil inti pengembang arsitektur digital dan pengawasan workflow AIFL</p>
              </div>
            </div>

            <div class="admin-team-grid">
              {{#each this.teamMembers as |member|}}
                <div class="admin-team-card">
                  <div class="team-avatar">{{member.avatar}}</div>
                  <h4 class="team-name">{{member.name}}</h4>
                  <div class="team-role">{{member.role}}</div>
                  <div class="team-email">📧 {{member.email}}</div>
                  <div class="team-specialty">⚡ {{member.specialty}}</div>
                </div>
              {{/each}}
            </div>
          </div>
        {{/if}}
      </main>
    </div>
  </template>
}
