package icu.aimerny.elix.registry;

import icu.aimerny.elix.Elix;
import icu.aimerny.elix.block.entity.ElixBlockEntity;
import icu.aimerny.elix.consts.IdConst;
import net.fabricmc.fabric.api.object.builder.v1.block.entity.FabricBlockEntityTypeBuilder;
import net.minecraft.block.entity.BlockEntityType;
import net.minecraft.registry.Registries;
import net.minecraft.registry.Registry;
import net.minecraft.util.Identifier;

public class ModBlockEntity {

    public static final BlockEntityType<ElixBlockEntity> ELIX_BLOCK_ENTITY = Registry.register(
            Registries.BLOCK_ENTITY_TYPE,
            new Identifier(Elix.MOD_ID, IdConst.ELIX_BLOCK_ENTITY),
            FabricBlockEntityTypeBuilder.create(ElixBlockEntity::new, ModBlock.ELIX_BLOCK).build()
    );

    public static void init() {

    }
}
