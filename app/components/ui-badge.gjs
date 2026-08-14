import Component from '@glimmer/component';

export default class UiBadge extends Component {
  get variantClass() {
    const variant = this.args.variant || 'fill';
    return `ui-badge-${variant}`;
  }

  get sizeClass() {
    const size = this.args.size;
    return size ? `ui-badge-${size}` : '';
  }

  <template>
    <div
      class="ui-badge {{this.variantClass}} {{this.sizeClass}}"
      ...attributes
    >
      {{#if @showDot}}
        <span class="ui-badge-dot"></span>
      {{/if}}
      <span class="ui-badge-text">
        {{yield}}
      </span>
    </div>
  </template>
}
