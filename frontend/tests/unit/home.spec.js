import { shallowMount, mount } from '@vue/test-utils'
import Home from '@/components/Home.vue'
import { RouterLinkStub } from '@vue/test-utils';


describe('Home.vue', () => {
  it('renders Home Component With Correct Div', () => {
    const wrapper = mount(Home,{
      stubs: {
        RouterLink: RouterLinkStub
      }

    })
    const helloDiv = wrapper.find('.hello')
    expect(helloDiv).toBeTruthy()
  })
})

describe('Home.vue', () => {
  it('The Div also includes the P tag', () => {
    const wrapper = mount(Home,{
      stubs: {
        RouterLink: RouterLinkStub
      }

    })
    const helloDiv = wrapper.find('.hello')
    const pTag = helloDiv.find('p');
    expect(pTag).toBeTruthy()
  })
})
