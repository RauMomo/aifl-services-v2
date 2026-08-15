import Component from '@glimmer/component';

export default class PortalAccessTemplate extends Component {
  <template>
    <div class="portal-route-content">
      {{outlet}}
    </div>
  </template>
}
