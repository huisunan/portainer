import { PortainerNamespace } from '@/react/kubernetes/namespaces/types';

import { name } from './name';
import { namespace } from './namespace';
import { created } from './created';

export function getColumns(namespaces?: PortainerNamespace[]) {
  return [name(namespaces), namespace, created];
}
