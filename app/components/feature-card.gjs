import Component from '@glimmer/component';

export default class FeatureCard extends Component {
  <template>
    <div class="feature-card">
      <div class="feature-icon">
        {{@icon}}
      </div>
      <h3 class="feature-title">{{@title}}</h3>
      <p class="feature-desc">{{@description}}</p>
    </div>
  </template>
}
